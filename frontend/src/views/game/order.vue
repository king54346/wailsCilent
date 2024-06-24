<template>
  <div style="margin:8px;">
    <el-row :gutter="15">
      <el-col :span="24">
        <el-card shadow="never" style="margin-bottom: 15px">
          <div class="welTop">
            <div class="icon"></div>
            <div class="main">
            <el-input
              v-model="search"
              placeholder="Search"
              @input="filterOrders"
              style="width: 20%;"
            ></el-input>
              <el-table :data="orderList" style="width: 100%; margin-top: 20px">
                <el-table-column prop="id" label="ID" width="50"></el-table-column>
                <el-table-column prop="phone" label="手机号"></el-table-column>
                <el-table-column prop="machine" label="机型"></el-table-column>
                <el-table-column prop="sdcard" label="内存卡容量(GB)"></el-table-column>
                <el-table-column prop="status" label="状态"></el-table-column>
                  <el-table-column label="操作">
                    <template #default="{ row }">
                      <div class="action-buttons">
                        <el-button @click="openStartOrderDialog(row.id)" size="mini">开始</el-button>
                        <el-button type="primary" @click="cancelOrderTask(row.id)" size="mini">取消</el-button>
                        <el-button @click="viewDetails(row.id)" type="primary" size="mini">详情</el-button>
                        <el-button @click="deleteOrder(row.id)" type="danger" size="mini">删除</el-button>
                      </div>
                    </template>
                  </el-table-column>
              </el-table>
               <el-pagination
                  background
                  layout="prev, pager, next"
                  :total="totalCount"
                  :page-size="pageSize"
                  @current-change="handlePageChange"
                  style="margin-top: 20px;"
                >
                </el-pagination>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
    <!-- Start Order Dialog -->
        <el-dialog title="开始安装" v-model="startOrderDialogVisible" width="80%">
          <el-form :model="startOrderForm">
            <el-form-item label="选择设备">
              <el-select v-model="startOrderForm.deviceId" placeholder="请选择设备">
                <el-option v-for="device in deviceList" :key="device.ID" :label="`名称: ${device.Name}, 序列号: ${device.Serial}`" :value="device.ID"></el-option>
              </el-select>
            </el-form-item>
          </el-form>
          <span slot="footer" class="dialog-footer">
            <el-button @click="startOrderDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="startOrder">开始</el-button>
          </span>
        </el-dialog>
    <!-- Details Dialog -->
        <el-dialog title="Order Details" v-model="detailsDialogVisible" width="80%">
          <el-table :data="orderDetails" style="width: 100%">
            <el-table-column prop="id" label="ID" width="50"></el-table-column>
            <el-table-column prop="name" label="游戏名"></el-table-column>
            <el-table-column prop="size" label="大小(GB)"></el-table-column>
            <el-table-column prop="status" label="状态"></el-table-column>
            <el-table-column prop="error_msg" label="错误信息"></el-table-column>
          </el-table>
        </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { ElMessage } from 'element-plus';
import { GetWorkOrder, DeleteWorkOrder,GetOrderGameList, GetDeviceList, AddTask,CancelTask } from '../../../wailsjs/go/main/App'

const orderList = ref([]);
const totalCount = ref(0);
const pageSize = ref(10);
const currentPage = ref(1);
const search = ref('')

const detailsDialogVisible = ref(false);
const orderDetails = ref([]);

const startOrderDialogVisible = ref(false);
const deviceList = ref([]);
const startOrderForm = reactive({
  deviceId: null,
  orderId: null,
});

const filterOrders = () => {
    fetchOrderList(currentPage.value,search.value)
}

const fetchOrderList = async (page = 1,search = "") => {
  try {
    const response = await GetWorkOrder({ "pageNo": page, "pageSize": pageSize.value, "search": search, "status": 0 });
    let result = response.result
    totalCount.value = result.totalCount;
    pageSize.value = result.pageSize;
    currentPage.value = result.pageNo;
    orderList.value = result.data;
  } catch (error) {
    console.error('Error fetching order list:', error);
    ElMessage.error('Error fetching order list');
  }
};

const openStartOrderDialog = async (orderId) => {
  startOrderForm.orderId = orderId;
  try {
    const response = await GetDeviceList();
    deviceList.value = response;
    startOrderDialogVisible.value = true;
  } catch (error) {
    console.error('Error fetching device list:', error);
    ElMessage.error('Error fetching device list');
  }
};

const cancelOrderTask = async (orderId) => {
try {
    let response = await CancelTask(orderId);
    if(response.code == 200){
      console.log(response.code)
      fetchOrderList(currentPage.value);
      ElMessage.success('Order started successfully');
    }else{
      ElMessage.error(''+response.msg);
    }
  } catch (error) {
    console.error('Error starting order:', error);
    ElMessage.error('Error starting order');
  }
};

const startOrder = async () => {
  try {
    await AddTask(startOrderForm.orderId,startOrderForm.deviceId);
    startOrderDialogVisible.value = false;
    fetchOrderList(currentPage.value);
    ElMessage.success('Order started successfully');
  } catch (error) {
    console.error('Error starting order:', error);
    ElMessage.error('Error starting order');
  }
};

const viewDetails =  async (orderId) => {
    const response = await GetOrderGameList(orderId);
    orderDetails.value = response;
    detailsDialogVisible.value = true;
  console.log('Viewing details for order:', orderId);
};

const deleteOrder = async (orderId) => {
  try {
    await DeleteWorkOrder([orderId]);
    fetchOrderList();
    ElMessage.success('Order deleted successfully');
  } catch (error) {
    console.error('Error deleting order:', error);
    ElMessage.error('Error deleting order');
  }
};

const handlePageChange = (page) => {
  currentPage.value = page;
  fetchOrderList(page);
};

onMounted(() => {
  fetchOrderList();
});
</script>

<style scoped>
.el-tag+.el-tag {
  margin-left: 10px;
}

.welTop {

}

.welTop .main {
  margin-left: 20px;
}

.welTop .main h2 {
  font-size: 20px;
  color: #3c4a54;
}

.welTop .main p {
  color: #999;
  margin-top: 10px;
  line-height: 1.8;
}

.welTop .icons {
  margin-left: auto;
  text-align: center;
}

.welTop .icons p {
  font-size: 12px;
}

.avatar-list .avatar {
  margin-left: -10px;
  border: 3px solid #fff;
  cursor: pointer;
}

.data-box .el-card {
  margin-bottom: 15px;
}

.data-box .item-background {
  background: #409eff;
  color: #fff;
}

.data-box .item-background .item h2 {
  color: #fff;
}

.data-box .item-background .item p {
  color: rgba(255, 255, 255, 0.5);
}

.data-box .item-background .item .icon i {
  background: rgba(255, 255, 255, 0.2);
}

.data-box .item {
  display: flex;
}

.data-box .item h2 {
  font-size: 12px;
  color: #999;
  font-weight: normal;
}

.data-box .item h4 {
  font-size: 25px;
  margin: 5px 0 5px 0;
}

.data-box .item p {
  font-size: 12px;
  color: #999;
}

.data-box .item .icon {
  margin-left: auto;
  display: flex;
  align-items: center;
  margin-right: 10px;
}

.data-box .item .icon i {
  font-size: 18px;
  background: #409eff;
  color: #fff;
  border-radius: 50%;
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.data-box .item p span.up {
  color: #f56c6c;
}

.data-box .item p span.down {
  color: #67c23a;
}
.action-buttons {
  display: flex;
  flex-wrap: wrap; /* 允许按钮组在小单元格中换行 */
  gap: 5px; /* 减小按钮之间的间距 */
  justify-content: center; /* 居中按钮 */
}
.el-button {
  margin: 0 2px; /* 调整每个按钮的间距 */
  padding: 0 5px; /* 减小按钮的内部填充 */
  font-size: 12px; /* 减小按钮的字体大小 */
}
</style>
