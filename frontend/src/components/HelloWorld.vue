<script setup>
import { GetGameList, AddGame, UpdateGame, DeleteGame, GetAnnouncement, UpdateAnnouncement } from '../../wailsjs/go/main/App'
import { GetWorkOrder, DeleteWorkOrder } from '../../wailsjs/go/main/App'
import { GetDeviceList } from '../../wailsjs/go/main/App'

function getAnnouncement() {
  GetAnnouncement().then(result => {
    for (let i = 0; i < result.data.length; i++) {
      console.log(result.data[i])
    }
  })
}

function updateAnnouncement() {
  UpdateAnnouncement(["这是公告1", "这是公告2", "这是公告3", "这是公告4"]).then(result => {
    console.log(result)
  })
}

function getGameList() {
  GetGameList().then(result => {
    // result is a json list of games
    // console log it to see the structure for each game
    for (let i = 0; i < result.length; i++) {
      console.log(result[i])
    }
  })
}

function addGames() {
  AddGame([
    {
      "name": "test1",
      "size": 512,
      "tags": ["tag1", "tag2"],
      "filepath": "/home/test1"
    }
  ]).then(result => {
    console.log(result)
  })
}

function updateGame() {
  UpdateGame({
    "id": 8,
    "name": "test1",
    "size": 103,
    "tags": [
      "tag1",
      "tag3"
    ],
    "filepath": "/home/test10086"
  }).then(result => {
    console.log(result)
  })
}

function deleteGame() {
  DeleteGame([8]).then(result => {
    console.log(result)
  })
}

function getWorkOrder() {
  GetWorkOrder({ "pageNo": 1, "pageSize": 10, "search": "", "status": 0 }).then(result => {
    if ("code" in result) {
      console.log("error!")
      console.log(result.msg)
    } else {
      console.log(result.result)
    }
  })
}

function deleteWorkOrder() {
  DeleteWorkOrder([2]).then(result => {
    console.log(result)
  })
}

function getDeviceList() {
  GetDeviceList().then(result => {
    console.log(result)
  })
}

</script>

<template>
  <main>
    <div class="input-box">
      <span class="result">公告: </span>
      <button class="btn" @click="getAnnouncement">Get</button>
      <button class="btn" @click="updateAnnouncement">Update</button>
    </div>
    <br>
    <div id="input" class="input-box">
      <span class="result">游戏: </span>
      <button class="btn" @click="getGameList">Get</button>
      <button class="btn" @click="addGames">Add</button>
      <button class="btn" @click="updateGame">Update</button>
      <button class="btn" @click="deleteGame">Delete</button>
    </div>
    <br>
    <div class="input-box">
      <span class="result">工单: </span>
      <button class="btn" @click="getWorkOrder">Get</button>
      <button class="btn" @click="deleteWorkOrder">Delete</button>
    </div>
    <br>
    <div class="input-box">
      <span class="result">MTP: </span>
      <button class="btn" @click="getDeviceList">Get Devices</button>
      <button class="btn" @click="sendFile">Send</button>
    </div>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
