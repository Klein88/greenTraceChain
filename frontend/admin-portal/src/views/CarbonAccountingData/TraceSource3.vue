<template>
  <div class="traceSource3" style="display: flex; flex-direction: column; margin-top: 30px; background-color: rgba(255, 255, 255, 0.5)">
    <div style="display: flex; flex-direction: row">
      <el-button type="primary" icon="ArrowLeft" @click="back(data)" style="width: 120px; margin-left: 20px">返回</el-button>
      <el-button type="default" icon="DArrowLeft" @click="quit" style="width: 120px; margin-left: 20px">退出</el-button>
    </div>
    <div style="display: flex; align-items: center; margin-top: 40px">
      <div style="width: 60px; height: 60px; display: flex; justify-content: center; align-items: center; background: #0d89ec; border-radius: 5px; margin-left: 50px">
        <svg t="1711801555895" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="16415" width="40" height="40"><path d="M512 96c229.76 0 416 186.24 416 416 0 30.762667-3.349333 61.098667-9.898667 90.624l-2.624 11.029333-62.08-15.573333c7.04-27.904 10.602667-56.746667 10.602667-86.08 0-194.410667-157.589333-352-352-352S160 317.589333 160 512 317.589333 864 512 864a350.506667 350.506667 0 0 0 225.408-81.621333l2.304-1.984-250.090667-245.546667 44.757334-45.717333 258.538666 253.824 38.314667 37.696-22.4 22.826666A414.762667 414.762667 0 0 1 512 928C282.24 928 96 741.76 96 512S282.24 96 512 96z m0 213.333333a202.666667 202.666667 0 0 1 198.613333 243.178667l-1.770666 7.893333-62.165334-15.232a138.666667 138.666667 0 1 0-107.968 102.933334l6.506667-1.429334 15.253333 62.144A202.666667 202.666667 0 1 1 512 309.333333z" fill="#ffffff" p-id="16416"></path></svg>
      </div>
      <div style="display: flex; flex-direction: column; justify-content: space-between; margin-left: 20px">
        <span>溯源编号</span>
        <span style="margin-top: 8px">{{ data.Txid }}</span>
      </div>
    </div>
    <div style="display: flex; justify-content: center">
      <el-steps :active="5" align-center style="width: 96%; margin-left: 25px; margin-top: 50px">
        <el-step>
          <template #title>
            <el-card style="margin: 40px 8px 0 8px; width: 180px">
              <template #header>
                <div style="display: flex; align-items: center">
                  <svg style="margin-left: 5px" t="1711803399668" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="20160" width="35" height="35"><path d="M360.597 926.64H210.89c-45.646 0-82.782-37.136-82.782-82.782V180.891c0-45.646 37.136-82.782 82.782-82.782h537.403c45.646 0 82.781 37.136 82.781 82.782v193.206c0 17.673-14.327 32-32 32s-32-14.327-32-32V180.891c0-10.181-8.601-18.782-18.781-18.782H210.89c-10.181 0-18.782 8.601-18.782 18.782v662.966c0 10.181 8.601 18.782 18.782 18.782h149.707c17.673 0 32 14.327 32 32s-14.327 32.001-32 32.001z" fill="#515151" p-id="20161"></path><path d="M618.585 324.35H292.35c-17.673 0-32-14.327-32-32s14.327-32 32-32h326.235c17.673 0 32 14.327 32 32s-14.327 32-32 32zM574.587 481.094H292.601c-17.673 0-32-14.327-32-32s14.327-32 32-32h281.986c17.673 0 32 14.327 32 32s-14.327 32-32 32zM482.592 641.085h-188.99c-17.673 0-32-14.327-32-32s14.327-32 32-32h188.99c17.673 0 32 14.327 32 32s-14.327 32-32 32zM453.311 921.976l-3.156-146.626c-5.508-17.409 2.855-35.333 12.959-46.674l263.139-295.372c7.36-8.262 17.431-13.394 28.355-14.45 11.676-1.129 23.359 2.65 32.029 10.374l108.471 96.634c17.425 15.523 18.743 42.595 2.954 60.387l-66.067 74.886c-11.692 13.254-31.915 14.519-45.166 2.826-13.253-11.691-14.519-31.914-2.826-45.166l52.461-59.464-76.253-67.932-246.373 276.553a60.84 60.84 0 0 1 0.272 4.618l1.675 77.806 68.047-7.82a81.528 81.528 0 0 0 3.246-3.442l0.226-0.251 135.027-148.398c11.894-13.07 32.131-14.026 45.205-2.133 13.071 11.895 14.026 32.133 2.133 45.205L634.747 881.818c-3.284 3.678-8.309 8.94-13.809 13.396-12.6 10.207-23.772 12.293-32.53 11.236l-135.097 15.526z m-2.105-143.727a0.203 0.203 0 0 0 0.011 0.026c-0.005-0.008-0.008-0.017-0.011-0.026z" fill="#515151" p-id="20162"></path><path d="M803.003 671.965c-8.355 0-16.701-3.252-22.977-9.725l-98.864-101.965c-12.302-12.688-11.989-32.947 0.699-45.25 12.688-12.3 32.946-11.989 45.249 0.699l98.864 101.965c12.302 12.688 11.989 32.947-0.699 45.25-6.215 6.027-14.248 9.026-22.272 9.026z" fill="#515151" p-id="20163"></path></svg>
                  <span style="margin-left: 10px">申请</span>
                </div>
              </template>
              <div style="display: flex; flex-direction: column">
                <span style="color: #909399">环节开始时间：</span>
                <span style="font-size: 14px">2024-01-10 15:00:33</span>
                <span style="margin-top: 10px; color: #909399">环节结束时间：</span>
                <span style="font-size: 14px">2024-01-10 15:00:33</span>
                <span style="margin-top: 10px; color: #909399">环节上链时间：</span>
                <span style="font-size: 14px">2024-01-10 15:00:33</span>
              </div>
              <el-button type="primary" style="margin-top: 15px">溯源详情</el-button>
            </el-card>
          </template>
        </el-step>
        <el-step>
          <template #title>
            <el-card style="margin-top: 40px; width: 180px; margin-left: 8px; margin-right: 8px">
              <template #header>
                <div style="display: flex; align-items: center">
                  <svg style="margin-left: 5px" t="1711803895051" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="25028" width="35" height="35"><path d="M63.715485 864.1424V159.858738a96.142115 96.142115 0 0 1 96.142115-96.142115h639.999289a96.142115 96.142115 0 0 1 96.142115 96.142115v319.7152h64.284373V159.858738A159.8576 159.8576 0 0 0 799.856889 0.001138H159.8576A159.8576 159.8576 0 0 0 0 159.858738v704.283662A159.8576 159.8576 0 0 0 159.8576 1024h304.355218v-63.715485H159.8576a96.142115 96.142115 0 0 1-96.142115-96.142115z" fill="#515151" p-id="25029"></path><path d="M1014.89665 649.102639l-127.999858-127.999858a33.564407 33.564407 0 0 0-45.51106 0l-320.284089 320.284089a33.564407 33.564407 0 0 0-9.102212 22.75553v127.999858a31.857742 31.857742 0 0 0 31.857742 31.857742h127.999858a33.564407 33.564407 0 0 0 22.755531-9.102212l176.924247-177.493136a31.857742 31.857742 0 1 0-44.942172-44.942173l-167.822036 167.822036h-83.057685v-83.057685l288.426346-287.857458 82.488797 82.488797-32.995519 33.564407-46.648837-43.804396a32.426631 32.426631 0 0 0-43.804395 47.217726l69.404367 64.284373a31.857742 31.857742 0 0 0 44.373284 0l77.937691-77.368803a33.564407 33.564407 0 0 0 0-46.648837zM159.8576 640.000427a31.857742 31.857742 0 0 0-31.857742 31.857742 32.426631 32.426631 0 0 0 31.857742 32.426631h383.999573a32.426631 32.426631 0 0 0 31.857743-32.426631 31.857742 31.857742 0 0 0-31.857743-31.857742zM787.341347 465.351732a31.857742 31.857742 0 0 0 31.857743-31.857743V256.000853h-179.199801a31.857742 31.857742 0 1 0 0 63.715485h78.506579L505.74166 501.191692l-159.288712-152.462053a31.857742 31.857742 0 0 0-44.373284 0l-170.666477 170.666477a32.426631 32.426631 0 0 0 45.511061 45.511061l147.910947-148.479835L482.417242 568.889395a31.288854 31.288854 0 0 0 42.666619 0l227.555303-195.697561v60.302155a31.857742 31.857742 0 0 0 34.702183 31.857743z" fill="#515151" p-id="25030"></path></svg>
                  <span style="margin-left: 10px">评估</span>
                </div>
              </template>
              <div style="display: flex; flex-direction: column">
                <span style="color: #909399">环节开始时间：</span>
                <span style="font-size: 14px">2024-01-10 15:00:33</span>
                <span style="margin-top: 10px; color: #909399">环节结束时间：</span>
                <span style="font-size: 14px">2024-01-10 15:00:39</span>
                <span style="margin-top: 10px; color: #909399">环节上链时间：</span>
                <span style="font-size: 14px">2024-01-10 15:00:39</span>
              </div>
              <el-button type="primary" style="margin-top: 15px">溯源详情</el-button>
            </el-card>
          </template>
        </el-step>
        <el-step>
          <template #title>
            <el-card style="margin-top: 40px; width: 180px; margin-left: 8px; margin-right: 8px">
              <template #header>
                <div style="display: flex; align-items: center">
                  <svg style="margin-left: 5px" t="1711804000114" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="26152" width="35" height="35"><path d="M896 960H128a32 32 0 0 1-32-32v-576A32 32 0 0 1 128 320h128a32 32 0 0 1 0 64H160v512h704V384H768a32 32 0 0 1 0-64h128a32 32 0 0 1 32 32v576a32 32 0 0 1-32 32z" fill="#515151" p-id="26153"></path><path d="M736 608h-448a32 32 0 0 1 0-64h448a32 32 0 0 1 0 64zM736 800h-448a32 32 0 0 1 0-64h448a32 32 0 0 1 0 64zM512 480a32 32 0 0 1-32-32V128a32 32 0 0 1 64 0v320a32 32 0 0 1-32 32z" fill="#515151" p-id="26154"></path><path d="M640 288a30.08 30.08 0 0 1-22.4-9.6L512 173.44 406.4 278.4a31.36 31.36 0 1 1-44.8-44.8l128-128a30.72 30.72 0 0 1 44.8 0l128 128a30.72 30.72 0 0 1 0 44.8 30.08 30.08 0 0 1-22.4 9.6z" fill="#515151" p-id="26155"></path></svg>
                  <span style="margin-left: 10px">发布</span>
                </div>
              </template>
              <div style="display: flex; flex-direction: column">
                <span style="color: #909399">环节开始时间：</span>
                <span style="font-size: 14px">2024-01-10 15:02:22</span>
                <span style="margin-top: 10px; color: #909399">环节结束时间：</span>
                <span style="font-size: 14px">2024-01-10 15:01:40</span>
                <span style="margin-top: 10px; color: #909399">环节上链时间：</span>
                <span style="font-size: 14px">2024-01-10 15:01:40</span>
              </div>
              <el-button type="primary" style="margin-top: 15px">溯源详情</el-button>
            </el-card>
          </template>
        </el-step>
        <el-step>
          <template #title>
            <el-card style="margin-top: 40px; width: 180px; margin-left: 8px; margin-right: 8px">
              <template #header>
                <div style="display: flex; align-items: center">
                  <svg style="margin-left: 5px" t="1711804350259" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="34177" width="35" height="35"><path d="M836.8 556.8H635.2l8.4-28.4c0-0.9 0.1-4.3 8.5-14.7 4.2-5.1 9.1-10.5 15.7-17.5l0.4-0.4c6.1-6.6 13.4-14.3 20.6-22.9 31.2-36.9 62.8-89 62.8-175.7 0-84.8-29-145.4-76.4-183.8C629.3 76.2 571 64 519.1 64c-51.9 0-112.2 12.2-160.6 48.7-50.1 37.8-82.7 98.3-82.7 184.6s31.5 141.4 62.6 179.6c6.9 8.4 14.1 16.6 21.5 24.5 1.5 1.7 3 3.3 4.5 4.8 4.7 5.1 8.5 9.3 12 13.3 3 3.3 5.7 6.9 8 10.8 0 9 4.5 20.4 4.3 26.4H187.2c-32.6 0.1-66.9 22.4-67 55.1l-0.1 176.4c0 20.5 16.6 37.1 37.1 37.1H867c9.8 0 19.3-3.9 26.2-10.8 7-7 10.9-16.4 10.9-26.2V611.9c-0.1-32.7-34.6-55.1-67.3-55.1z m0 201.6H187.2V624h201.6c12.9 0.3 37.9-1.2 47.9-9.3 9.2-7.9 15.6-18.5 18.3-30.3 3.7-14.5 3.7-34 3.7-51.9v-4.1c0-25.8-16-45.6-26-57.1-4.4-5.3-9.6-10.8-14.5-16l-3.7-4c-6.2-6.8-12.4-13.6-18.4-21.1-23.3-28.5-46.1-67.5-46.1-132.9 0-65.5 23.6-103 53.3-125.5h0.1c31.4-23.7 74.1-33.8 115.9-33.8 41.7 0 81.2 10.1 109.5 33 27 21.8 49 59.4 49 126.2 0 65-22.5 100.8-45.4 127.7-6 7.2-12.2 13.8-18.6 20.6l-0.3 0.3c-6.1 6.5-12.7 13.7-18.7 20.9-11.8 14.4-25.2 34.7-25.2 61.6v0.6c0 13.3 0 24.6 0.3 33.5 0.1 8.9 1.1 17.8 3 26.5 3 12.9 11.1 24 22.4 30.8 11.7 6.8 35.3 4.1 40 4.1h201.6v134.6zM867 892.8H157.1c-20.5 0-37.1 15-37.1 33.6 0 8.9 3.9 17.5 10.9 23.8 7 6.3 16.4 9.8 26.2 9.8H867c20.5 0 37-15 37-33.6 0-18.5-16.6-33.6-37-33.6z" fill="#515151" p-id="34178"></path></svg>
                  <span style="margin-left: 10px">签约</span>
                </div>
              </template>
              <div style="display: flex; flex-direction: column">
                <span style="color: #909399">环节开始时间：</span>
                <span style="font-size: 14px">2024-01-10 15:02:22</span>
                <span style="margin-top: 10px; color: #909399">环节结束时间：</span>
                <span style="font-size: 14px">2024-01-10 15:02:51</span>
                <span style="margin-top: 10px; color: #909399">环节上链时间：</span>
                <span style="font-size: 14px">2024-01-10 15:02:51</span>
              </div>
              <el-button type="primary" style="margin-top: 15px">溯源详情</el-button>
            </el-card>
          </template>
        </el-step>
        <el-step>
          <template #title>
            <el-card style="margin-top: 40px; width: 180px; margin-left: 8px; margin-right: 8px">
              <template #header>
                <div style="display: flex; align-items: center">
                  <svg style="margin-left: 5px" t="1711804442077" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="35286" width="37" height="37"><path d="M511.24419 97.52381l9.728 7.875047 10.24 7.801905 6.339048 4.681143 7.216762 5.12 12.239238 8.387047 9.094095 5.973334 9.801143 6.217143 10.48381 6.41219 11.166476 6.534095 11.776 6.607238 6.119619 3.31581 12.678095 6.607238 13.287619 6.534095c56.368762 27.014095 126.878476 50.712381 202.776381 52.46781l9.142857 0.121905v279.503238c-2.169905 162.279619-79.043048 263.92381-159.280762 325.973333l-10.483809 7.850667-5.241905 3.779047-10.459428 7.192381-10.410667 6.802286c-3.462095 2.194286-6.92419 4.315429-10.337524 6.363428l-10.24 5.973334-10.044952 5.607619-9.874286 5.193143-9.630476 4.827428-4.729905 2.267429-13.750857 6.314667-8.752762 3.754666-12.434286 5.046857-7.777524 2.974477-10.800761 3.900952-9.654858 3.218286-7.387428 2.31619-6.363429 1.877333-7.411809 2.023619L512 926.47619l-10.142476-2.438095-11.824762-3.34019-8.923429-2.779429-10.166857-3.388952-7.411809-2.608762-7.875048-2.925714-12.55619-4.973715-8.850286-3.705904a534.674286 534.674286 0 0 1-4.559238-1.999239l-9.337905-4.242285-9.606095-4.583619-9.849905-4.973715-10.069333-5.339428-5.095619-2.80381-10.288762-5.924571a494.713905 494.713905 0 0 1-15.676953-9.654857l-10.50819-6.948572-10.556953-7.387428-10.556952-7.801905c-78.969905-60.416-154.136381-158.96381-157.379048-316.099048L170.666667 512V232.277333c77.628952 0 148.821333-22.625524 205.848381-49.347047l13.409523-6.485334 12.824381-6.582857 6.192762-3.291428 11.946667-6.631619 11.288381-6.582858 10.678857-6.485333 9.97181-6.339047 9.264761-6.095239 8.557715-5.851428 11.434666-8.118857 9.654858-7.192381 11.946666-9.411048 7.558095-6.339047z m0.48762 91.550476l-4.924953 3.462095a726.454857 726.454857 0 0 1-59.66019 36.498286c-61.464381 33.694476-126.098286 57.977905-192.926477 69.290666l-10.410666 1.633524V512c0 122.148571 48.859429 209.310476 132.973714 271.701333l6.89981 4.998096a440.758857 440.758857 0 0 0 120.978285 59.977142l6.972953 2.096762 4.827428-1.462857 5.802667-1.950476c144.505905-48.932571 251.757714-150.503619 257.706667-326.363429l0.219428-9.313523v-211.968l-10.727619-1.706667c-59.367619-10.24-117.564952-30.915048-173.787428-59.416381l-11.215239-5.802667a761.246476 761.246476 0 0 1-60.732952-35.718095l-11.995428-7.996952z m145.212952 172.470857l51.443809 51.98019-261.851428 259.218286-131.047619-130.998857 51.687619-51.736381 79.62819 79.579429 210.139429-208.042667z" p-id="35287" fill="#515151"></path></svg>
                  <span style="margin-left: 10px">认证</span>
                </div>
              </template>
              <div style="display: flex; flex-direction: column">
                <span style="color: #909399">环节开始时间：</span>
                <span style="font-size: 14px">2024-01-10 15:02:51</span>
                <span style="margin-top: 10px; color: #909399">环节结束时间：</span>
                <span style="font-size: 14px">2024-01-10 15:03:02</span>
                <span style="margin-top: 10px; color: #909399">环节上链时间：</span>
                <span style="font-size: 14px">2024-01-10 15:03:02</span>
              </div>
              <el-button type="primary" style="margin-top: 15px" @click="traceSource(data)">溯源详情</el-button>
            </el-card>
          </template>
        </el-step>
      </el-steps>
    </div>
  </div>
</template>

<script setup>
import {useRoute, useRouter} from "vue-router";
import {onMounted, ref} from "vue";

const route = useRoute()
const router = useRouter()

const back = (data) => {
  // 在页面跳转时保存数据到 sessionStorage 中
  // sessionStorage.setItem('myData3', JSON.stringify(data));
  sessionStorage.removeItem('myData3');
  router.push('/traceSource1/'+data.CompanyId)
}

const quit = () => {
  sessionStorage.removeItem('myData1');
  sessionStorage.removeItem('myData2');
  sessionStorage.removeItem('myData3');
  sessionStorage.removeItem('myData4');
  router.push('/carbonAccountingData')
}

const Id = route.query.Id
const CompanyId = route.query.CompanyId
const Txid = route.query.Txid
const data = ref({})

onMounted(()=>{
  data.value.Id = Id
  data.value.CompanyId = CompanyId
  data.value.Txid = Txid
  console.log(Id)

  // 在页面加载时从 sessionStorage 中获取数据并显示在页面上
  const savedData = getDataFromSessionStorage();
  if (savedData) {
    data.value = savedData;
  }
})

const traceSource = (data) => {
  console.log(data)
  // 在页面跳转时保存数据到 sessionStorage 中
  sessionStorage.setItem('myData3', JSON.stringify(data));
  router.push({
    path: '/traceSource4/'+data.CompanyId,
    query: {
      Id: data.Id,
      CompanyId: data.CompanyId
    }
  })
}

// 在页面加载时从 sessionStorage 中获取数据
const getDataFromSessionStorage = () => {
  const data = sessionStorage.getItem('myData3');
  return data ? JSON.parse(data) : null;
}
</script>

<style scoped>
/* 添加带有透明度的背景颜色 */
.traceSource3::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-image: url('@/assets/traceSource3.png');
  background-size: cover;
  z-index: -1;
  opacity: 0.3;
}
</style>