<template>
  <div class="login">
    <video src="../images/climatecoin_video_home.mp4" class="video-bg" autoplay muted loop playsinline style="position: fixed; object-fit: fill; width: 100%; height: 100vh; top: 0px; left: 0px" @click="toggleLoginCardVisibility"></video>
    <div class="animated-card" v-if="showLoginCard">
      <el-card class="box-card" style="background: white">
        <img src="@/assets/logo.png" alt="logo" height="50px" width="320px">
        <el-divider content-position="center">登录</el-divider>
        <div>
          <el-form ref="loginFormRef" :model="User">
            <el-form-item label="用户名">
              <el-input type="text" v-model="User.UserID" placeholder="请输入用户名"></el-input>
            </el-form-item>
            <el-form-item label="密码">
              <el-input type="password" v-model="User.Password" placeholder="请输入密码" show-password style="margin-left: 13px"></el-input>
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
            <el-button type="primary" text @click="toggleRegisterCardVisibility">注册账户</el-button>
          </div>
        </template>
      </el-card>
    </div>
    <div class="animated-card" v-if="showRegisterCard">
      <el-card class="box-card">
        <img src="@/assets/logo.png" alt="logo" height="50px" width="320px">
        <el-divider content-position="center">注册</el-divider>
        <div>
          <el-form :model="User1">
            <el-form-item label="用户名">
              <el-input type="text" v-model="User1.UserID" placeholder="请输入用户名"></el-input>
            </el-form-item>
            <el-form-item label="密码">
              <el-input type="password" v-model="User1.Password" placeholder="请输入密码" show-password style="margin-left: 13px"></el-input>
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

<script setup>
import avatar from "../assets/user.png"
import {getCurrentInstance, ref} from "vue";
import {useRouter} from "vue-router";
import {ElMessage} from "element-plus";

const avatarUrl = ref(avatar)

const router = useRouter()
const currentInstance = getCurrentInstance()
const { proxy } = currentInstance

//表单绑定的响应式对象
const User = ref({
  UserID: '',
  Password: ''
})

const User1 = ref({
  UserID: '',
  Password: ''
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

//表单的引用对象
const loginFormRef = ref()

const onSubmit = () => {
  //在这里进行用户名和密码的验证
  //校验表单
  loginFormRef.value.validate((valid)=>{
    if (valid){
      proxy.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=CompanyLogin", {
        CompanyID: User.value.UserID,
        Password: User.value.Password
      }).then(res=>{
        console.log(res.data)
        if (res.data.msg == 'success'){
          switch (res.data.data.Type){
            case 3: router.push('/home')
          }
          ElMessage.success('登录成功')
          sessionStorage.setItem('UserID',User.value.UserID)
        } else {
          ElMessage.error('用户名和密码不匹配')
        }
      })
    }
  })
}

const onSubmitRegister = async () => {
  await proxy.$axios.post('http://47.97.176.174:8087/chainmaker?cmb=DAPRegisiter', {
    'CompanyID' : User1.value.UserID,
    'Password' : User1.value.Password
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

<style scoped>
.login {
  background-size: cover;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

.video-bg {
  z-index: 0
}

.card-header {
  display: flex;
  justify-content: space-around;
  align-items: center;
}

.box-card {
  z-index: 1;
  background: white;
  width: 360px;
  border-radius: 10px;
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

.animated-card {
  animation: fadeInBlur 2s ease-in-out forwards;
  transform: translate(0%, 0%);
}
</style>
