<template>
  <div style="margin:8px;">
    <el-row :gutter="15">
      <el-col :span="24">
        <el-card shadow="never" style="margin-bottom: 15px">
          <div class="welTop">
            <div class="icon"></div>
            <div class="main">
              <div class="header">
                    <el-button
                      type="primary"
                      @click="openAddDialog"
                    >
                      添加
                    </el-button>
                    <el-input
                      v-model="search"
                      placeholder="Search"
                      @input="filterGames"
                      style="width: 20%;"
                    ></el-input>
                  </div>
              <el-table :data="gameList" style="width: 100%; margin-top: 20px">
                <el-table-column prop="id" label="ID" width="50"></el-table-column>
                <el-table-column prop="name" label="游戏名"></el-table-column>
                <el-table-column prop="size" label="大小(GB)"></el-table-column>
                <el-table-column prop="tags" label="标签">
                  <template #default="{ row }">
                    <el-tag v-for="tag in row.tags" :key="tag">{{ tag }}</el-tag>
                  </template>
                </el-table-column>
                <el-table-column prop="filepath" label="文件路径"></el-table-column>
                <el-table-column label="操作" width="180">
                  <template #default="{ row }">
                    <el-button @click="openEditDialog(row)" size="mini">编辑</el-button>
                    <el-button @click="deleteGame(row.id)" type="danger" size="mini">删除</el-button>
                  </template>
                </el-table-column>
              </el-table>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <el-dialog title="添加" v-model="addDialogVisible">
      <el-form :model="form">
        <el-form-item label="游戏名">
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="大小">
          <el-input v-model.number="form.size" type="number"></el-input>
        </el-form-item>
               <el-form-item label="标签">
                 <el-select
                   v-model="form.tags"
                   multiple
                   filterable
                   allow-create
                   default-first-option
                   :reserve-keyword="false"
                 >

                   <el-option v-for="tag in availableTags" :key="tag" :label="tag" :value="tag"></el-option>
                 </el-select>
               </el-form-item>
            <el-form-item label="文件路径">
              <el-input v-model="form.filepath"></el-input>
            </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="addDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="addGame">添加</el-button>
      </div>
    </el-dialog>

    <el-dialog title="编辑" v-model="editDialogVisible">
      <el-form :model="form">
        <el-form-item label="游戏名">
          <el-input v-model="form.name"></el-input>
        </el-form-item>
        <el-form-item label="大小">
          <el-input v-model.number="form.size" type="number"></el-input>
        </el-form-item>
             <el-form-item label="标签">
                         <el-select
                           v-model="form.tags"
                           multiple
                           filterable
                           allow-create
                           default-first-option
                           :reserve-keyword="false"
                         >
                           <el-option v-for="tag in availableTags" :key="tag" :label="tag" :value="tag"></el-option>
                         </el-select>
                       </el-form-item>
                <el-form-item label="文件路径">
                  <el-input v-model="form.filepath"></el-input>
                </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="editDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="updateGame">保存</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { ElMessage } from 'element-plus';
import { GetGameList, AddGame, UpdateGame, DeleteGame, GetAnnouncement, UpdateAnnouncement } from '../../../wailsjs/go/main/App'

const gameList = ref([]);
const addDialogVisible = ref(false);
const editDialogVisible = ref(false);
const form = reactive({
  id: null,
  name: '',
  size: 0,
  tags: [],
  filepath: ''
});
const availableTags = ref([]);

const search = ref('')

const filterGames = () => {
    fetchGameList(search.value)
}


const fetchGameList = async (search="") => {
  try {
    const response = await GetGameList(search);
    gameList.value = response;
  } catch (error) {
    console.error('Error fetching game list:', error);
    ElMessage.error('Error fetching game list');
  }
};


const openAddDialog = () => {
  form.id = null;
  form.name = '';
  form.size = 0;
  form.tags = [];
  form.filepath = '';
  addDialogVisible.value = true;
};
const handleSizeInput = (event) => {
  form.size = parseFloat(event.target.value);
};

const addGame = async () => {
  try {
    console.log(form)
    await AddGame([form]);
    addDialogVisible.value = false;
    fetchGameList();
    ElMessage.success('Game added successfully');
  } catch (error) {
    console.error('Error adding game:', error);
    ElMessage.error('Error adding game');
  }
};

const openEditDialog = (game) => {
  form.id = game.id;
  form.name = game.name;
  form.size = game.size;
  form.tags = game.tags;
  form.filepath = game.filepath;
  editDialogVisible.value = true;
};

const updateGame = async () => {
  try {
    await UpdateGame(form);
    editDialogVisible.value = false;
    fetchGameList();
    ElMessage.success('Game updated successfully');
  } catch (error) {
    console.error('Error updating game:', error);
    ElMessage.error('Error updating game');
  }
};


const deleteGame = async (gameId) => {
  try {
    await DeleteGame([gameId]);
    fetchGameList();
    ElMessage.success('Game deleted successfully');
  } catch (error) {
    console.error('Error deleting game:', error);
    ElMessage.error('Error deleting game');
  }
};


onMounted(() => {
  fetchGameList();
});
</script>

<style  scoped>
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
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}
</style>
