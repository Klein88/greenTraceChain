<script setup>
import avatar from "@/assets/login/avatar.jpg"
import {ref} from "vue";
import {useRouter} from "vue-router";
import axios from "axios";
import { ElMessage } from 'element-plus'

const avatarUrl = ref(avatar)

const router = useRouter()

const User = ref({
  username: 'DAPAdmin1632',
  password: 'admin1632'
})

const User1 = ref({
  username: '',
  password: ''
})

const showLoginCard = ref(false);
const showRegisterCard = ref(false);


const toggleLoginCardVisibility = () => {
  showLoginCard.value = true

  showRegisterCard.value = false;
};

const toggleRegisterCardVisibility = () => {
  showLoginCard.value = false; // 隐藏登录卡片
  showRegisterCard.value = true; // 显示注册卡片
};

/*
const onSubmit = () => {
  router.push({
    path: '/Home'
  })
}
*/

const onSubmit = async () => {
  await axios.post('http://47.97.176.174:8087/chainmaker?cmb=CompanyLogin', {
    'CompanyID' : User.value.username,
    'Password' : User.value.password
  })
      .then(function (response) {
        if (response.data.msg === "请前往注册") {
          router.push({
            path: '/'
          })
          ElMessage.error('登录失败')
        } else {
          switch (response.data.data.Type) {
            case 0:
              router.push({
                path: '/Home'
              })
              ElMessage({
                message: '登录成功',
                type: 'success',
              })
              break;
            case 1:
              router.push({
                path: '/Home'
              })
              break;
            case 2:
              router.push({
                path: '/Home'
              })
              break;
            case 3:
              router.push({
                path: '/Home'
              })
              break;
            default :
              router.push({
                path: '/'
              })
              ElMessage.error('登录失败')
          }
        }
        console.log(response)
      })
      .catch(function (error) {
        console.log(error);
      })
}

const onSubmitRegister = async () => {
  await axios.post('http://47.97.176.174:8087/chainmaker?cmb=DAPRegisiter', {
    'CompanyID' : User1.value.username,
    'Password' : User1.value.password
  })
      .then(function (response) {
        ElMessage({
          message: '注册成功',
          type: 'success',
        })
        console.log(response)
      })
      .catch(function (error) {
        console.log(error);
      })
}
</script>

<template>
  <div class="login">
    <video src="@/assets/login/login.mp4"  autoplay loop muted @click="toggleLoginCardVisibility"></video>
    <div class="animated-card" v-if="showLoginCard">
      <el-card class="box-card">
        <template #header>
          <img src="@/assets/home/logo.png" alt="logo" height="50px" width="320px">
          <el-divider content-position="center">登录</el-divider>
          <div class="card-header">
            <el-avatar :src="avatarUrl"></el-avatar>
            <el-button disabled>数据审核员</el-button>
          </div>
        </template>
        <div>
          <el-form :model="User">
            <el-form-item label="用户名">
              <el-input type="text" v-model="User.username" placeholder="请输入用户名"></el-input>
            </el-form-item>
            <el-form-item label="密码">
              <el-input type="password" v-model="User.password" placeholder="请输入密码" show-password></el-input>
            </el-form-item>
            <div style="display: flex; justify-content: space-between; align-items: center">
              <el-checkbox label="自动登录" />
              <el-button type="primary" text>忘记密码</el-button>
            </div>
            <div style="display: flex; justify-content: center; align-items: center">
              <el-button type="primary" @click="onSubmit">立即登录</el-button>
            </div>
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
            <el-button type="primary" text @click="toggleRegisterCardVisibility">注册账户</el-button>
          </div>
        </template>
      </el-card>
    </div>
    <div class="animated-card" v-if="showRegisterCard">
      <el-card class="box-card">
        <template #header>
          <img src="@/assets/home/logo.png" alt="logo" height="50px" width="320px">
          <el-divider content-position="center">注册</el-divider>
          <div class="card-header">
            <el-avatar :src="avatarUrl"></el-avatar>
            <el-button disabled>数据审核员</el-button>
          </div>
        </template>
        <div>
          <el-form :model="User1">
            <el-form-item label="用户名">
              <el-input type="text" v-model="User1.username" placeholder="请输入用户名"></el-input>
            </el-form-item>
            <el-form-item label="密码">
              <el-input type="password" v-model="User1.password" placeholder="请输入密码" show-password></el-input>
            </el-form-item>
            <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 10px">
              <el-checkbox label="我已阅读并同意《用户注册协议》" />
            </div>
            <div style="display: flex; justify-content: center; align-items: center">
              <el-button type="primary" @click="onSubmitRegister">立即注册</el-button>
            </div>
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
            <el-button type="primary" text @click="toggleLoginCardVisibility">返回登录</el-button>
          </div>
        </template>
      </el-card>
    </div>
  </div>
</template>

<style scoped>
.login {
  //background-image: url("../assets/login/login.jpg");
  //background-size: cover;
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
  z-index: 1;
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

video {
  height: 100vh;
  width: 100vw;
  position: fixed;
  object-fit: fill;
}

.animated-card {
  animation: fadeInBlur 2s ease-in-out forwards;
  transform: translate(-50%, -50%);
}

@keyframes fadeInBlur {
  0% {
    opacity: 0;
    filter: blur(10px);
    transform: scale(0.95);
  }
  100% {
    opacity: 1;
    filter: blur(0);
    transform: scale(1);
  }
}
</style>
