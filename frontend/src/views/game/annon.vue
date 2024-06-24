<template>
  <div style="margin:8px;">
    <el-row :gutter="15">
      <el-col :span="24">
        <el-card shadow="never" style="margin-bottom: 15px">
          <div class="welTop">
            <div class="icon">
            </div>
            <div class="main">
               <el-button type="primary" @click="fetchAnnouncement">刷新公告</el-button>
               <el-button type="success" @click="addAnnouncement">添加公告</el-button>
               <div class="announcement">
                           <div v-if="announcements.length">
                             <h3>公告列表:</h3>
                             <ul>
                               <li v-for="(announcement, index) in announcements" :key="index">
                                 <div v-if="!editingIndex.includes(index)">
                                   {{ announcement }}
                                   <el-button type="text" @click="editAnnouncement(index)">编辑</el-button>
                                    <el-button type="text" @click="deleteAnnouncement(index)">删除</el-button>
                                 </div>
                                 <div v-else>
                                   <el-input v-model="announcementEdits[index]" placeholder="Edit announcement"></el-input>
                                   <el-button type="primary" @click="saveAnnouncement(index)">保存</el-button>
                                   <el-button type="danger" @click="cancelEdit(index)">取消</el-button>
                                 </div>
                               </li>
                             </ul>
                           </div>
                 </div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { GetAnnouncement, UpdateAnnouncement } from '../../../wailsjs/go/main/App'
import { ref, onMounted } from 'vue';

const announcements = ref([]);
const editingIndex = ref([]);
const announcementEdits = ref([]);

const fetchAnnouncement = async () => {
  try {
    const result = await GetAnnouncement();
    if (result.code === 0 && result.msg === 'ok') {
      announcements.value = result.data;
      announcementEdits.value = [...result.data];
      console.log(result);
    } else {
      console.error('Failed to fetch announcements');
    }
  } catch (error) {
    console.error('Error fetching announcements:', error);
  }
};
const deleteAnnouncement = (index) => {
  announcements.value.splice(index, 1);
  announcementEdits.value.splice(index, 1);
  updateAnnouncement(announcements.value);
};

const updateAnnouncement = async (updatedAnnouncements) => {
  try {
    const result = await UpdateAnnouncement(updatedAnnouncements);
    if (result.code === 0 && result.msg === 'ok') {
      await fetchAnnouncement(); // Fetch updated announcements after updating
      console.log(result);
    } else {
      console.error('Failed to update announcements');
    }
  } catch (error) {
    console.error('Error updating announcements:', error);
  }
};

const editAnnouncement = (index) => {
  editingIndex.value.push(index);
};

const saveAnnouncement = (index) => {
  announcements.value[index] = announcementEdits.value[index];
  updateAnnouncement(announcements.value);
  editingIndex.value = editingIndex.value.filter(i => i !== index);
};

const cancelEdit = (index) => {
  announcementEdits.value[index] = announcements.value[index];
  editingIndex.value = editingIndex.value.filter(i => i !== index);
};

const addAnnouncement = () => {
  announcements.value.push("New Announcement");
  announcementEdits.value.push("New Announcement");
  editAnnouncement(announcements.value.length - 1);
};

// Fetch announcements on component mount
onMounted(fetchAnnouncement);
</script>


<style  scoped>
.el-tag+.el-tag {
  margin-left: 10px;
}

.welTop {
  display: flex;
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
.announcement {
  padding: 20px;
}

.announcement h3 {
  margin-top: 20px;
}

.announcement ul {
  list-style-type: none;
  padding: 0;
}

.announcement li {
  background: #f4f4f4;
  margin: 5px 0;
  padding: 10px;
  border-radius: 4px;
}
</style>
