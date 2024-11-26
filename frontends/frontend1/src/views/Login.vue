<script setup>
import {ref} from "vue";
import {useRouter} from "vue-router";
import axios from "axios";
import {ElMessage} from "element-plus";
import { useStore } from 'vuex';

const avatarUrl = ref(avatar)
const store=useStore()
const router = useRouter()

const User = ref({
  username: 'admin',
  password: 'admin'
})

const onSubmit = () => {
  store.dispatch('login',{
    companyid: User.value.username,
    password: User.value.password
  }).then(() => {
    router.push({
      path: '/DashBoard'
    })
  }).catch((error) => {
    console.error('登录失败:', error);
    ElMessage.error('登录失败，请检查用户名和密码');
  });
}
const reg =()=>{
  router.push({ name: 'Regisiter' });
}

/*
const onSubmit = async () => {
  const res = await axios.get('路径', {
    params: {
      'username': newContent.value.name,
      'password': newContent.value.password,
    }
  })
      .then(function (response) {
        router.push({
          path: '/Dashboard'
        })
        console.log(response)
      })
      .catch(function (error) {
        console.log(error);
      })
}
*/
</script>

<template>
  <div class="login">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <el-avatar :src="avatarUrl"></el-avatar>
          <el-button class="button" text>用户端</el-button>
        </div>
      </template>
      <div>
        <el-form :model="User">
          <el-form-item label="用户ID">
            <el-input type="text" v-model="User.username" placeholder="请输入用户ID"></el-input>
          </el-form-item>
          <el-form-item label="密码">
            <el-input type="password" v-model="User.password" placeholder="请输入密码" show-password></el-input>
          </el-form-item>
          <div style="display: flex; justify-content: space-between; align-items: center">
            <el-checkbox label="自动登录" />
            <el-button type="primary" text>忘记密码</el-button>
          </div>
          <div style="display: flex; justify-content: center; align-items: center">
          <el-button type="primary" @click="onSubmit">立即登录</el-button></div>
        </el-form>
      </div>
      <template #footer>
        <div class="card-footer">
          <div>
            <el-button text disabled>其他登录方式</el-button>
            <a href="https://auth.alipay.com/login/index.htm?goto=https%3A%2F%2Fwww.alipay.com%2F">
              <img src="../assets/login/zfb.svg" alt="支付宝">
            </a>
            <a href="https://login.taobao.com/member/login.jhtml?spm=a21bo.jianhua.754894437.1.5af92a89ZPVLGq&f=top&redirectURL=https%3A%2F%2Fwww.taobao.com%2F">
              <img src="../assets/login/tb.svg" alt="淘宝">
            </a>
          </div>
          <el-button type="primary" @click="reg" text>注册账户</el-button>
        </div>
      </template>
    </el-card>
  </div>
</template>

<style scoped>
.login {
  background-image: url("../assets/login/login.jpg");
  background-size: cover;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

.card-header {
  display: flex;
  justify-content: space-around;
  align-items: center;
}

.box-card {
  width: 360px;
  border-radius: 0;
  border: 0;
}

.card-footer {
  display: flex;
  justify-content: space-around;
  align-items: center;
}

.card-footer img {
  vertical-align: middle;
  margin: 5px;
}
</style>
