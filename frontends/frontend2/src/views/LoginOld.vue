<script setup>
import avatar from "@/assets/public/avatar.jpg";
import {ref} from "vue";
import {useRouter} from "vue-router";
import axios from "axios";
import { User, Lock } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'

const avatarUrl = ref(avatar);
const router = useRouter();

const Login = ref({
  username: '',
  password: ''
});
const Register = ref({
  username: '',
  password: ''
});

const showLoginCard = ref(false);
const showRegisterCard = ref(false);

const toggleLoginCardVisibility = () => {
  showLoginCard.value = true;
  showRegisterCard.value = false;
};
const toggleRegisterCardVisibility = () => {
  showLoginCard.value = false;
  showRegisterCard.value = true;
};

const onSubmitLogin = async () => {
  await axios.post('http://47.97.176.174:8087/chainmaker?cmb=CompanyLogin', {
    CompanyID : Login.value.username,
    Password : Login.value.password
  })
      .then(function (response) {
        if (response.data.msg === "请前往注册") {
          ElMessage({
            message: '请前往注册',
            type: 'warning',
          })
        } else {
          router.push({
            path: '/Home'
          })
          ElMessage({
            message: '登录成功',
            type: 'success',
          })
        }
        console.log(response)
      })
      .catch(function (error) {
        ElMessage.error('登录失败')
        console.log(error)
      })
}
const onSubmitRegister = async () => {
  await axios.post('http://47.97.176.174:8087/chainmaker?cmb=DAPRegisiter', {
    CompanyID : Register.value.username,
    Password : Register.value.password
  })
      .then(function (response) {
        ElMessage({
          message: '注册成功',
          type: 'success',
        })
        console.log(response)
      })
      .catch(function (error) {
        ElMessage.error('注册失败')
        console.log(error)
      })
}
</script>

<template>
  <div class="login">
    <video src="@/assets/login/login.mp4"  autoplay loop muted @click="toggleLoginCardVisibility"></video>
    <div class="animated-card" v-if="showLoginCard">
      <el-card class="box-card">
        <template #header>
          <img src="../assets/public/logo.png" alt="logo" height="50px" width="320px">
          <el-divider content-position="center">登录</el-divider>
          <div class="card-header">
            <el-avatar :src="avatarUrl"></el-avatar>
            <el-button disabled>数据审核员</el-button>
          </div>
        </template>
        <div>
          <el-form :model="Login">
            <el-form-item>
              <el-input type="text" v-model="Login.username" size="large" placeholder="请输入用户名" :prefix-icon="User"></el-input>
            </el-form-item>
            <el-form-item>
              <el-input type="password" v-model="Login.password" size="large" placeholder="请输入密码" show-password :prefix-icon="Lock"></el-input>
            </el-form-item>
            <div style="display: flex; justify-content: space-between; align-items: center;  margin-bottom: 15px;">
              <div>
                <el-checkbox label="自动登录" />
                <el-checkbox label="记住密码" />
              </div>
              <el-button type="primary" text>忘记密码</el-button>
            </div>
            <div style="display: flex; justify-content: center; align-items: center;">
              <el-button type="primary" @click="onSubmitLogin">立即登录</el-button>
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
          <img src="../assets/public/logo.png" alt="logo" height="50px" width="320px">
          <el-divider content-position="center">注册</el-divider>
          <div class="card-header">
            <el-avatar :src="avatarUrl"></el-avatar>
            <el-button disabled>数据审核员</el-button>
          </div>
        </template>
        <div>
          <el-form :model="Register">
            <el-form-item>
              <el-input type="text" v-model="Register.username" size="large" placeholder="请输入用户名" :prefix-icon="User"></el-input>
            </el-form-item>
            <el-form-item>
              <el-input type="password" v-model="Register.password" size="large" placeholder="请输入密码" show-password :prefix-icon="Lock"></el-input>
            </el-form-item>
            <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 15px;">
              <el-checkbox>
                我已阅读并同意
                <span style="color: #409EFF;">《用户注册协议》</span>
              </el-checkbox>
            </div>
            <div style="display: flex; justify-content: center; align-items: center;">
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
  /*
  background-image: url("../assets/login/login.jpg");
  background-size: cover;
  */
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

video {
  height: 100vh;
  width: 100vw;
  position: fixed;
  object-fit: cover;
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
