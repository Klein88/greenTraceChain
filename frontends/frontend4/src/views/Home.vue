<template>
  <div class="common-layout">
    <el-container>
      <el-aside class="aside">
        <el-card class="aside-card">
          <div class="aside-card-logo">
            <img src="../assets/logo.png" alt="logo">
          </div>
          <el-scrollbar>
            <el-menu
                class="el-menu-vertical-demo"
                @open="handleOpen"
                @close="handleClose"
                router
                @select="handleSelect"
                background-color="#283342"
                text-color="#909399"
                active-text-color="#FFFFFF"
            >
              <el-menu-item index="/home">
                <el-icon><House /></el-icon>
                <span>控制台</span>
              </el-menu-item>
              <el-sub-menu index="2">
                <template #title>
                  <el-icon><Aim/></el-icon>
                  <span>用户审核</span>
                </template>
                <el-menu-item index="/transactionData">
                  <el-icon><Money /></el-icon>
                  <span>交易数据</span>
                </el-menu-item>
                <el-menu-item index="/carbonAccountingData">
                  <el-icon><DocumentCopy /></el-icon>
                  <span>碳核算数据</span>
                </el-menu-item>
              </el-sub-menu>
              <el-sub-menu index="3">
                <template #title>
                  <el-icon>
                    <svg t="1709788484664" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1697" width="20" height="20"><path d="M781.2 510.9l62.4 44.3-173.2 247.4-62.7 21.7-1.1-63.9 174.6-249.5m-5.7-68.9c-11.7 0-23.3 5.3-30.1 15.4L553.7 731.2c-3.7 5.5-7.3 14.6-7.3 21.9l1.8 104.1c0 21.4 18 36.4 36.7 36.4 3.6 0 7.3-0.6 10.8-1.7l100.4-34.7c7.3-1.8 14.6-7.3 18.3-12.8l191.7-273.8c11-16.4 7.3-38.3-9.1-51.1l-100.4-71.2c-6.4-4.2-13.8-6.3-21.1-6.3zM415 111c-125.9 0-228 102.1-228 228 0 81.7 43 153.4 107.6 193.6-91.1 42-153.7 131.9-171.6 244.4l-8.5 67.5C110 880.3 137.9 912 174 912h269.6c1.1 0 2.1-0.1 3.2-0.2 1.1 0.1 2.1 0.2 3.2 0.2 16.6 0 30-13.4 30-30s-13.4-30-30-30c-1.1 0-2.2 0.1-3.2 0.2-1-0.1-2.1-0.2-3.2-0.2H208c-18.1 0-32-15.8-29.8-33.7l4.1-32.7c5.3-32.5 14.9-62.7 28.6-90 13.4-26.6 30.4-49.7 50.5-68.9 41.2-39 94.2-59.7 153.5-59.7 30 0 58.8 5.5 85.5 16.5 5.9 2.4 11.7 5.1 17.3 8 4.4 2.3 9.2 3.4 13.9 3.4 9 0 17.8-4 23.6-11.5 11.5-14.8 6.9-36.5-9.8-45.1-3.5-1.8-7.1-3.6-10.8-5.3C599.7 492.8 643 421 643 339c0-125.9-102.1-228-228-228z m0 396c-44.9 0-87.1-17.5-118.8-49.2S247 383.9 247 339s17.5-87.1 49.2-118.8S370.1 171 415 171c44.9 0 87.1 17.5 118.8 49.2C565.5 251.9 583 294.1 583 339s-17.5 87.1-49.2 118.8C502.1 489.5 459.9 507 415 507z" p-id="1698"></path></svg>
                  </el-icon>
                  <span>用户管理</span>
                </template>
                <el-menu-item index="/userRequest">
                  <el-icon><EditPen /></el-icon>
                  <span>用户请求</span>
                </el-menu-item>
                <el-menu-item index="/userView">
                  <el-icon><Postcard /></el-icon>
                  <span>用户查看</span>
                </el-menu-item>
              </el-sub-menu>
              <el-menu-item index="/contractManagement">
                <el-icon><Memo /></el-icon>
                <span>合约管理</span>
              </el-menu-item>
              <el-menu-item index="/contractWriting">
                <el-icon><DocumentAdd /></el-icon>
                <span>合约撰写</span>
              </el-menu-item>
<!--              <el-menu-item index="/chainManagement">-->
<!--                <el-icon><Connection /></el-icon>-->
<!--                <span>链上管理</span>-->
<!--              </el-menu-item>-->
            </el-menu>
          </el-scrollbar>
        </el-card>
      </el-aside>
      <el-container>
        <el-header class="header">
          <el-card shadow="never" class="header-card">
            <div class="toolbar">
              <el-input suffix-icon="Search" placeholder="搜索..." style="width: 280px"></el-input>
              <el-button style="background: #6ebe71; color: white; margin-left: 55px; width: 230px; height: 40px; border-radius: 10px" @click="chain">
                <svg t="1712040655422" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="21825" width="20" height="20"><path d="M512 23.272727C242.082909 23.272727 23.272727 242.082909 23.272727 512s218.810182 488.727273 488.727273 488.727273 488.727273-218.810182 488.727273-488.727273S781.917091 23.272727 512 23.272727m0 69.818182c230.981818 0 418.909091 187.927273 418.909091 418.909091s-187.927273 418.909091-418.909091 418.909091S93.090909 742.981818 93.090909 512 281.018182 93.090909 512 93.090909m-116.340364 421.213091a115.968 115.968 0 0 1 34.071273-84.573091 115.898182 115.898182 0 0 1 84.573091-34.071273l228.072727-114.036363-114.036363 228.072727a115.968 115.968 0 0 1-34.071273 84.573091 115.898182 115.898182 0 0 1-84.596364 34.071273l-228.072727 114.036363 114.059636-228.072727zM512 465.454545a46.568727 46.568727 0 0 0 0 93.09091 46.568727 46.568727 0 0 0 0-93.09091z" fill="#ffffff" p-id="21826"></path></svg>
                <span style="margin-left: 15px">进入区块链浏览器</span>
              </el-button>
              <el-button icon="Notification" style="background: #6EB5FF; color: white; margin-left: 55px; width: 150px; height: 40px; border-radius: 10px" @click="drawer = true">查看事项</el-button>
              <el-icon :size="23" style="margin-left: 50px;"><Bell /></el-icon>
              <img :src="avatarUrl" alt="avatar" style="height: 45px; width: 45px; border-radius: 50%; margin-left: 50px">
              <div style="display: flex; flex-direction: column; margin-left: 15px">
                <span style="font-size: 18px; font-weight: bold">{{ UserID }}</span>
                <span style="font-size: 13px; margin-top: 5px">管理员</span>
              </div>
            </div>
          </el-card>
        </el-header>
        <el-main class="main1">
          <div v-if="currentRouteIsHome">
            <div style="display: flex; flex-direction: column">
              <div style="display: flex; justify-content: center; align-items: center;">
                <el-card style="width: 21%; height: 250px; background: rgba(24,255,228,0.17); border-radius: 10px; margin-top: 30px">
                  <div style="display: flex; flex-direction: column">
                    <div style="width: 90px; height: 90px; background: white; border-radius: 50%; display: flex; justify-content: center; align-items: center">
                      <svg t="1711518187449" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="10486" width="60" height="60"><path d="M928 832 896 832 896 384c0-70.4-57.6-128-128-128L576 256 576 192c0-70.4-57.6-128-128-128L192 64C121.6 64 64 121.6 64 192l0 640L32 832C14.08 832 0 846.08 0 864 0 881.92 14.08 896 32 896l896 0c17.92 0 32-14.08 32-32C960 846.08 945.92 832 928 832zM512 256l0 64 0 512L128 832 128 192c0-35.2 28.8-64 64-64l256 0c35.2 0 64 28.8 64 64L512 256zM832 832 576 832 576 320l192 0c35.2 0 64 28.8 64 64L832 832zM416 256l-192 0C206.08 256 192 270.08 192 288 192 305.92 206.08 320 224 320l192 0C433.92 320 448 305.92 448 288 448 270.08 433.92 256 416 256zM416 448l-192 0C206.08 448 192 462.08 192 480 192 497.92 206.08 512 224 512l192 0C433.92 512 448 497.92 448 480 448 462.08 433.92 448 416 448zM416 640l-192 0C206.08 640 192 654.08 192 672S206.08 704 224 704l192 0C433.92 704 448 689.92 448 672S433.92 640 416 640zM736 448l-64 0C654.08 448 640 462.08 640 480 640 497.92 654.08 512 672 512l64 0C753.92 512 768 497.92 768 480 768 462.08 753.92 448 736 448zM736 640l-64 0c-17.92 0-32 14.08-32 32s14.08 32 32 32l64 0c17.92 0 32-14.08 32-32S753.92 640 736 640z" p-id="10487" fill="#00BFFF"></path></svg>
                    </div>
                    <span style="font-size: 23px; font-weight: bold; margin-top: 8px">企业</span>
                    <span style="font-size: 21px; margin-top: 13px">{{ company1 }}</span>
                  </div>
                  <div style="display: flex; align-items: center; margin-top: 22px">
                    <span style="font-size: 12px; color: #8A2BE2">+20%</span>
                    <span style="font-size: 12px; margin-left: 6px; color: darkgray">This Month</span>
                  </div>
                </el-card>
                <el-card style="width: 21%; height: 250px; background: rgba(255,232,69,0.2); border-radius: 10px; margin-top: 30px; margin-left: 34px">
                  <div style="display: flex; flex-direction: column">
                    <div style="width: 90px; height: 90px; background: white; border-radius: 50%; display: flex; justify-content: center; align-items: center">
                      <svg t="1711518357179" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="12687" width="60" height="60"><path d="M296 249.984a8 8 0 0 0-8 8v48c0 4.416 3.584 8 8 8h384a8 8 0 0 0 8-8v-48a8 8 0 0 0-8-8h-384zM480 393.984H296a8 8 0 0 0-8 8v48c0 4.416 3.584 8 8 8H480a8 8 0 0 0 8-8v-48a8 8 0 0 0-8-8z m-48 458.048h-224v-704H768v320c0 4.352 3.584 8 8 8H832a8 8 0 0 0 8-8V108.032a32 32 0 0 0-32-32h-640a32 32 0 0 0-32 32v784a32 32 0 0 0 32 32H432a8 8 0 0 0 8-8.064v-55.936a8 8 0 0 0-8-8.064z m440-88.064h-144v-36.544a112 112 0 1 0-64 0v36.608h-144a16 16 0 0 0-16 16v151.936c0 8.832 7.168 16 16 16h352a16 16 0 0 0 16-16v-151.936a16 16 0 0 0-16-16z m-225.984-144a50.048 50.048 0 1 1 100.032 0.064 50.048 50.048 0 0 1-100.032 0z m179.968 266.048H566.016v-60.032h259.968v60.032z" fill="orange" p-id="12688"></path></svg>
                    </div>
                    <span style="font-size: 23px; font-weight: bold; margin-top: 8px">审核员</span>
                    <span style="font-size: 21px; margin-top: 13px">{{ company2 }}</span>
                  </div>
                  <div style="display: flex; align-items: center; margin-top: 22px">
                    <span style="font-size: 12px; color: #8A2BE2">+3%</span>
                    <span style="font-size: 12px; margin-left: 6px; color: darkgray">This Month</span>
                  </div>
                </el-card>
                <el-card style="width: 21%; height: 250px; background: rgba(255,168,178,0.2); border-radius: 10px; margin-top: 30px; margin-left: 34px">
                  <div style="display: flex; flex-direction: column">
                    <div style="width: 90px; height: 90px; background: white; border-radius: 50%; display: flex; justify-content: center; align-items: center">
                      <svg t="1711517939381" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="6663" width="60" height="60"><path d="M383.078 570.982c-4.403 25.063-4.403 88.32 0 110.234 8.55 43.136 34.714 84.608 67.303 109.875l15.36-122.189h62.72l13.517 122.19c28.57-18.433 50.534-40.04 68.89-109.876 12.364-47.206 9.574-79.232 5.452-109.9l29.44 1.382c70.502 3.993 150.554 13.388 184.064 30.438 36.557 18.611 45.286 125.491 78.362 297.805l-805.786-0.768 4.275-58.035c6.989-86.093 21.274-209.434 60.007-249.114 8.422-8.576 80.563-15.949 216.396-21.99v-0.052z m113.997 0l47.155 48.896-18.15 32.64h-58.112l-17.587-32.563 46.694-48.896v-0.077zM481.92 76.954c119.22 3.584 228.582 56.064 247.424 90.956 18.944 34.944-40.678 73.626-67.584 83.61l3.61 4.096c1.715 5.76 6.067 17.408 4.582 27.52 1.1 4.66-0.717 7.245 0 13.005l-8.14 18.176c3.583 195.072-138.574 223.053-168.96 223.053-30.337 0-194.612-28.032-185.166-229.53-5.504-6.81-8.37-11.7-8.37-11.7l-0.385-13.004c-0.153-11.008 0-25.958 1.997-31.437l-2.688-0.947c-11.98-4.506-60.058-25.165-60.902-66.202-0.973-45.44 117.043-107.033 232.601-107.75l11.981 0.154z m154.06 139.008H326.35v33.536H635.98v-33.536z m-143.743-104.55c-22.298 0-40.32 15.436-40.32 34.508l0.512 4.506c-19.917-1.588-36.506 0.332-36.557 0.332 31.104 12.084 44.826 20.813 44.826 20.813 7.321 5.683 16 7.86 23.756 8.474l5.581 0.256 2.1 0.179 5.375-0.384 8.09-1.536 0.768-0.384c3.277-1.024 6.4-2.458 9.344-4.224l52.736-23.091c-11.802 0.179-24.832 1.1-36.813 2.099 0.589-2.304 0.896-4.66 0.896-7.04 0-19.072-18.022-34.509-40.294-34.509z" p-id="6664" fill="#EE82EE"></path></svg>
                    </div>
                    <span style="font-size: 23px; font-weight: bold; margin-top: 8px">监管机构</span>
                    <span style="font-size: 21px; margin-top: 13px">{{ company3 }}</span>
                  </div>
                  <div style="display: flex; align-items: center; margin-top: 22px">
                    <span style="font-size: 12px; color: #8A2BE2">+1%</span>
                    <span style="font-size: 12px; margin-left: 6px; color: darkgray">This Month</span>
                  </div>
                </el-card>
                <el-card style="width: 21%; height: 250px; background: rgba(0,0,0,0.06); border-radius: 10px; margin-top: 30px; margin-left: 34px">
                  <div style="display: flex; flex-direction: column">
                    <div style="width: 90px; height: 90px; background: white; border-radius: 50%; display: flex; justify-content: center; align-items: center">
                      <svg t="1711520277222" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="20609" width="60" height="60"><path d="M1004.590899 728.576l-109.056-90.112-20.992-24.576-3.072-391.68c0-7.168-3.072-14.336-8.704-18.944l-160.256-142.336c-4.608-4.096-10.752-6.656-16.896-6.656H236.590899c-48.128 0-87.04 38.912-87.04 87.04V547.84c-28.16-9.216-59.392-9.728-90.112-1.024-14.848 4.096-28.16 12.288-38.4 24.064-30.208 33.792-27.648 85.504 6.144 115.712L261.678899 898.048c17.408 15.872 40.96 24.064 65.024 22.016l320-27.648 129.536 103.424c22.016 17.408 53.76 14.848 72.704-6.144l162.304-184.32c0.512-0.512 0.512-0.512 1.024-1.536 18.432-23.04 14.848-56.32-7.68-75.264zM697.902899 124.928l105.472 93.696h-105.472V124.928z m-461.312-18.944h410.112v138.752c0 14.336 11.264 25.6 25.6 25.6h148.992l2.048 283.648c-30.208-31.232-71.68-48.128-115.2-46.08L425.518899 519.68c-52.736 2.048-93.696 45.568-93.696 98.304 0 23.552 8.192 45.056 21.504 61.952l-21.504 2.048-131.072-104.96v-435.2c0-19.968 16.384-35.84 35.84-35.84z m735.232 664.576c-13.824 15.36-27.136 30.72-40.96 46.08-17.92 20.48-31.744 35.84-40.96 46.08-54.784 61.44-81.92 92.16-82.432 92.16L739.886899 901.12l-67.584-53.76c-6.656-5.632-15.36-8.192-24.064-7.168l-327.168 28.672c-9.216 1.024-18.432-2.56-25.6-8.704L61.998899 648.192c-12.288-11.264-13.312-29.696-2.048-41.984 3.584-4.096 8.704-7.168 15.872-9.216 31.232-8.704 64.512-2.048 89.6 17.92l150.016 120.32 227.84-19.968c14.336-1.024 28.16-5.632 40.96-12.8l45.056-25.6-201.216-11.776c-24.576-1.536-44.032-22.016-44.032-47.104 0-24.576 19.456-45.056 44.032-46.08l282.112-11.776c30.72-1.536 60.416 11.776 80.384 35.328l69.12 81.408 112.128 92.672c0.512-0.512 0.512 0.512 0 1.024z" p-id="20610" fill="#707070"></path><path d="M296.494899 436.736H722.478899c14.336 0 25.6-11.264 25.6-25.6s-11.264-25.6-25.6-25.6H296.494899c-14.336 0-25.6 11.264-25.6 25.6s11.264 25.6 25.6 25.6zM296.494899 290.816H512.558899c14.336 0 25.6-11.264 25.6-25.6s-11.264-25.6-25.6-25.6H296.494899c-14.336 0-25.6 11.264-25.6 25.6s11.264 25.6 25.6 25.6z" p-id="20611" fill="#707070"></path></svg>
                    </div>
                    <span style="font-size: 23px; font-weight: bold; margin-top: 8px">智能合约</span>
                    <span style="font-size: 21px; margin-top: 13px">11</span>
                  </div>
                  <div style="display: flex; align-items: center; margin-top: 22px">
                    <span style="font-size: 12px; color: #8A2BE2">+2%</span>
                    <span style="font-size: 12px; margin-left: 6px; color: darkgray">This Month</span>
                  </div>
                </el-card>
              </div>
              <div style="display: flex; justify-content: center; align-items: center;">
                <el-card style="width: 45%; height: 300px; border-radius: 8px; margin-top: 30px">
                  <div style="display: flex; align-items: center">
                    <svg t="1711725776037" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4255" width="40" height="40"><path d="M893.088 612.832c17.952 5.952 27.776 25.28 21.824 43.328a424.864 424.864 0 0 1-404.16 291.936c-197.248 0-363.584-134.816-411.552-317.28l-2.848 5.248a34.368 34.368 0 0 1-62.976-6.944 34.24 34.24 0 0 1 2.88-26.048l53.248-96.96a34.272 34.272 0 0 1 64.32 16.512c0 196.8 160.128 356.928 356.928 356.928 154.4 0 290.656-98.432 339.008-244.896a34.24 34.24 0 0 1 43.328-21.824z m78.528-218.56a34.24 34.24 0 0 1 13.76 46.432l-53.248 98.272a34.336 34.336 0 0 1-64.416-16.32c0-196.8-160.096-357.024-356.896-357.024A355.232 355.232 0 0 0 169.472 417.6a34.304 34.304 0 0 1-65.504-20.224A423.392 423.392 0 0 1 510.72 97.056c196.96 0 363.104 134.56 411.36 316.48l3.008-5.504a34.24 34.24 0 0 1 46.432-13.76h0.064z m-547.072 217.984c0 24.8 21.632 46.624 51.424 55.68V542.4c-53.568-21.952-120-49.216-120-119.104 0-62.144 51.552-114.176 120-126.176v-27.072a34.272 34.272 0 1 1 68.576 0v29.248c64.224 14.624 111.904 64.384 111.904 124a34.272 34.272 0 0 1-68.576 0c0-22.464-17.92-42.272-43.328-52.608v125.856c50.08 20.704 111.904 49.888 111.904 115.744 0 59.616-47.68 109.376-111.904 124.032v38.944a34.272 34.272 0 0 1-68.576 0v-36.8c-68.448-12-120-64-120-126.176a34.272 34.272 0 1 1 68.576 0z m0-188.992c0 17.92 19.104 30.24 51.424 44.64v-100.32c-29.792 9.088-51.424 30.816-51.424 55.68z m163.328 188.992c0-15.936-16.896-28.224-43.328-40.8v93.44c25.44-10.336 43.328-30.176 43.328-52.64z" fill="#3CB371" p-id="4256"></path></svg>
                    <span style="font-size: 20px; font-weight: bold; margin-left: 10px">碳币汇率</span>
                    <el-button icon="Refresh" @click="reloadPage" style="margin-left: 200px">点击查看</el-button>
                  </div>
                  <div ref="chartRef" style="width: 100%; height: 180px; margin-top: 18px"></div>
                </el-card>
                <el-card style="width: 45%; height: 300px; border-radius: 8px; margin-top: 30px; margin-left: 35px">
                  <div style="display: flex; align-items: center">
                    <svg t="1711725884228" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5276" width="40" height="40"><path d="M746.264 744.485v155.273c0 10.181-8.243 18.303-18.303 18.303H618.02c-10.182 0-18.303-8.243-18.303-18.303v-92.485c-33.697 11.151-57.818 17.333-57.818 17.333v95.636c0 36.728 13.94 51.03 50.667 51.03h153.09c28-0.12 50.667-22.666 50.91-50.666V716.364c-13.818 8-31.03 17.818-50.303 28.12z m-24.606-692H592.203c-36.727 0-50.667 39.757-50.667 76.606v561.576l-0.12 0.242c19.757-6.545 39.15-13.94 58.181-22.06V125.454c0-10.182 8.242-18.303 18.303-18.303h109.94c10.181 0 18.302 8.242 18.302 18.303v458.787c14.303-10.666 28.243-21.939 41.819-33.697l8.606-13.09v-433.94c0.12-36.727-38.182-51.03-74.91-51.03zM413.9 295.515H260.688c-13.576-0.121-26.667 5.212-36.121 14.788-9.576 9.576-14.91 22.424-14.91 36v396.606c22.546 0.243 40 0.364 57.819-0.242V374.06c0-10.788 8.727-19.637 19.636-19.637h101.455c10.788 0 19.636 8.728 19.636 19.637v353.454c38.546-7.273 56.849-14.182 56.849-14.182V346.545c-0.243-28.242-23.03-51.03-51.152-51.03z m-5.818 560v42.91c0 10.787-8.727 19.636-19.637 19.636H286.991c-10.788 0-19.636-8.728-19.636-19.637v-28.727c-1.697 0-3.516 0.121-5.213 0.121h-0.12v-0.606a380.97 380.97 0 0 1-52.607 0v63.758c0 10.182 4 20 11.152 27.151 7.272 7.152 16.97 11.152 27.151 11.152h178.91c10.181 0 20-4 27.15-11.152 7.152-7.273 11.152-16.97 11.152-27.151v-89.334c-18.788 4.606-37.818 8.485-56.848 11.88z" p-id="5277"></path><path d="M556.567 789.818c-94.788 32.849-194.182 50.667-294.667 52.606-118.182 0-204.97-32.363-226.182-94.182-21.212-61.939 29.455-141.818 125.576-213.697v85.334C109.173 668 92.204 708.242 98.627 727.515c8 24 62.788 49.697 163.273 49.697 93.09-2.303 185.455-19.03 273.333-49.697 101.819-32 196.485-83.394 278.91-151.273 56.12-50.06 74.424-92.242 67.757-111.757-4.606-13.455-24-27.637-57.333-37.333v-67.516c63.03 13.94 106.788 42.303 120.97 83.88 36.605 107.272-137.455 262.302-388.364 346.545l-0.606-0.243z" p-id="5278"></path></svg>
                    <span style="font-size: 20px; font-weight: bold; margin-left: 10px">企业类型</span>
                    <el-button icon="Refresh" @click="createPie" style="margin-left: 200px">点击查看</el-button>
                  </div>
                  <div id="pie" style="width: 100%; height: 250px; margin-top: 18px"></div>
                </el-card>
              </div>
              <div style="display: flex; justify-content: center; align-items: center; margin-bottom: 30px">
                <el-card style="width: 45%; border-radius: 8px; margin-top: 30px;">
                  <div style="display: flex; align-items: center">
                    <svg t="1711726839812" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="7193" width="40" height="40"><path d="M769.7 837.1H256.2c-43.2 0-78.3-35.1-78.3-78.3V356c0-43.2 35.1-78.3 78.3-78.3h513.5c43.2 0 78.3 35.1 78.3 78.3v402.8c-0.1 43.2-35.3 78.3-78.3 78.3zM256.2 327.7c-15.6 0-28.3 12.7-28.3 28.3v402.8c0 15.6 12.7 28.3 28.3 28.3h513.5c15.6 0 28.2-12.7 28.3-28.4V356c0-15.6-12.7-28.3-28.3-28.3H256.2zM757.9 189.3h-490c-13.8 0-25 11.2-25 25s11.2 25 25 25h490c13.8 0 25-11.2 25-25s-11.2-25-25-25z" p-id="7194" fill="#FF0000"></path><path d="M645.4 525.1h-78.3c-8.6 0-13.9-9.3-9.6-16.7l52.1-90.2c6.4-11 3.2-25.4-7.6-32.1-11.3-7.1-26.2-3.3-32.8 8.2l-59.1 102.4c-3 5.1-10.3 5.1-13.3 0l-59.1-102.4c-6.5-11.2-20.8-15.1-32.1-8.6-11.2 6.5-15.1 20.8-8.6 32.1l52.4 90.7c4.3 7.4-1.1 16.7-9.6 16.7h-78.3c-12.7 0-23.6 9.9-24 22.6-0.5 13.2 10.2 24.2 23.4 24.2H459c11.6 0 20.9 9.4 20.9 20.9v32.5s0 0.1-0.1 0.1h-91.2c-12.8 0-23.6 9.9-24.1 22.7-0.4 13.3 10.2 24.3 23.5 24.3h91.8s0.1 0 0.1 0.1v47.7c0 12.8 9.9 23.6 22.7 24.1 13.3 0.4 24.3-10.2 24.3-23.5v-48.3s0-0.1 0.1-0.1h91.2c12.8 0 23.6-9.9 24.1-22.7 0.4-13.3-10.2-24.3-23.5-24.3H527s-0.1 0-0.1-0.1V592c0-11 8.9-19.9 19.9-19.9h99.1c13.2 0 23.9-11 23.5-24.3-0.4-12.8-11.3-22.7-24-22.7z" p-id="7195" fill="#FF0000"></path></svg>
                    <span style="font-size: 18px; font-weight: bold; margin-left: 10px">交易信息</span>
                  </div>
                  <el-table :data="transaction" stripe style="width: 100%; margin-top: 15px; height: 200px">
                    <el-table-column prop="CompanyId" align="center" label="公司编号" width="120"></el-table-column>
                    <el-table-column label="碳币" prop="CarbonCoin" align="center" width="120"></el-table-column>
                    <el-table-column label="碳额度" prop="CarbonCredit" align="center" width="120"></el-table-column>
                    <el-table-column label="操作" align="center">
                      <template #default>
                        <router-link to="/transactionData">
                          <Button label="查看" size="small" class="mybutton"></Button>
                        </router-link>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
                <el-card style="width: 45%; border-radius: 8px; margin-left: 35px; margin-top: 30px;">
                  <div style="display: flex; align-items: center">
                    <svg t="1711770109347" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="8263" width="40" height="40"><path d="M874.666667 0a64 64 0 0 1 63.701333 57.856L938.666667 64v320.768a42.666667 42.666667 0 0 1-85.034667 4.992L853.333333 384.768V85.333333H170.666667v853.333334h299.050666a42.666667 42.666667 0 0 1 42.368 37.674666l0.298667 4.992a42.666667 42.666667 0 0 1-37.674667 42.368L469.717333 1024H149.333333a64 64 0 0 1-63.701333-57.856L85.333333 960v-896A64 64 0 0 1 143.189333 0.298667L149.333333 0h725.333334z" fill="#20B2AA" p-id="8264"></path><path d="M704 170.666667a64 64 0 0 1 63.701333 57.856L768 234.666667V341.333333a42.666667 42.666667 0 0 1-85.034667 4.992L682.666667 341.333333V256H341.333333v170.666667h213.76a42.666667 42.666667 0 0 1 42.368 37.674666l0.298667 4.992a42.666667 42.666667 0 0 1-37.717333 42.368L555.093333 512H320a64 64 0 0 1-63.701333-57.856L256 448v-213.333333a64 64 0 0 1 57.856-63.701334L320 170.666667h384z" fill="#20B2AA" p-id="8265"></path><path d="M256 682.666667m42.666667 0l170.666666 0q42.666667 0 42.666667 42.666666l0 0q0 42.666667-42.666667 42.666667l-170.666666 0q-42.666667 0-42.666667-42.666667l0 0q0-42.666667 42.666667-42.666666Z" fill="#20B2AA" p-id="8266"></path><path d="M426.666667 597.333333m0 42.666667l0 170.666667q0 42.666667-42.666667 42.666666l0 0q-42.666667 0-42.666667-42.666666l0-170.666667q0-42.666667 42.666667-42.666667l0 0q42.666667 0 42.666667 42.666667Z" fill="#20B2AA" p-id="8267"></path><path d="M810.666667 512c115.2 0 213.333333 43.605333 213.333333 118.528l-0.170667 5.461333L1024 640v256.256C1024 974.933333 925.44 1024 810.666667 1024c-111.36 0-207.36-46.165333-213.077334-120.661333L597.333333 896.256V640l0.213334-4.053333-0.213334-5.418667c0-72.661333 92.245333-115.84 202.922667-118.4L810.666667 512z m0 374.528c-47.445333 0-92.032-7.424-128-21.333333v31.061333c0 3.712 7.552 12.586667 26.368 21.973333 25.557333 12.757333 62.037333 20.437333 101.632 20.437334 39.594667 0 76.074667-7.68 101.632-20.437334 16.725333-8.32 24.533333-16.298667 26.069333-20.565333l0.298667-1.408 0.042666-31.104c-35.968 13.952-80.554667 21.333333-128.042666 21.333333z m128.042666-158.848c-35.968 13.952-80.554667 21.333333-128.042666 21.333333-47.445333 0-92.032-7.381333-128-21.333333V768l5.546666 4.181333c3.882667 2.730667 10.069333 6.442667 19.114667 10.453334 26.026667 11.562667 63.146667 18.56 103.338667 18.56 40.192 0 77.354667-6.997333 103.338666-18.56 18.090667-8.021333 24.661333-14.933333 24.661334-14.634667l0.042666-40.32zM810.666667 597.333333c-40.192 0-77.354667 6.997333-103.338667 18.56-7.253333 3.2-12.629333 6.229333-16.512 8.704l-2.602667 1.706667-5.546666 4.224 3.456 2.688c3.626667 2.730667 10.325333 7.125333 21.205333 11.946667 26.026667 11.52 63.146667 18.56 103.338667 18.56 40.192 0 77.354667-6.997333 103.338666-18.56 5.418667-2.389333 9.813333-4.693333 13.312-6.741334l5.802667-3.712 5.546667-4.181333-5.546667-4.224a117.76 117.76 0 0 0-19.114667-10.410667C887.978667 604.330667 850.858667 597.333333 810.666667 597.333333z" fill="#20B2AA" p-id="8268"></path></svg>
                    <span style="font-size: 18px; font-weight: bold; margin-left: 10px">碳核算信息</span>
                  </div>
                  <el-table :data="carbonAccounting" stripe style="width: 100%; margin-top: 15px; height: 200px">
                    <el-table-column prop="CompanyId" label="公司编号"align="center"></el-table-column>
                    <el-table-column prop="Type" label="企业类型" align="center">
                      <template #default="scope">
                        <el-tag :type="typeTagType(scope.row.Type)" :disable-transitions="false">
                          {{ typeTagText(scope.row.Type) }}
                        </el-tag>
                      </template>
                    </el-table-column>
                    <el-table-column label="操作" align="center">
                      <template #default>
                        <router-link to="/carbonAccountingData">
                          <Button label="查看" size="small" class="mybutton"></Button>
                        </router-link>
                      </template>
                    </el-table-column>
                  </el-table>
                </el-card>
              </div>
<!--              <img src="../assets/chain.jpeg" width="93%" height="200" @click="chain" style="margin: 0 10px 30px 40px; border-radius: 10px">-->
            </div>
            <el-drawer v-model="drawer" direction="rtl" :with-header="false">
              <el-card style="width: 100%; height: 180px; border-radius: 10px">
                <div style="display: flex; align-items: center">
                  <svg t="1711589242371" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="23543" width="40" height="40"><path d="M521.8304 198.912l-223.104 127.4624-98.2272 0.0256A57.6 57.6 0 0 0 142.8992 384v256l0.128 3.9424a57.6 57.6 0 0 0 57.472 53.6576l98.176-0.0256 223.1552 127.5136a57.6 57.6 0 0 0 86.1952-49.9968V248.9088a57.6 57.6 0 0 0-86.1952-49.9968z m22.1696 61.0304v504.0896l-228.2752-130.432-108.8512-0.0256v-243.2l108.8512 0.0256 228.2752-130.4576zM653.952 371.0208a32 32 0 0 1 44.8512-5.9648c42.1376 32.2304 63.3088 79.488 63.3088 138.624s-21.1712 106.3936-63.3088 138.624a32 32 0 0 1-38.8864-50.8416c25.5488-19.5328 38.1952-47.7696 38.1952-87.808 0-40.0128-12.6464-68.224-38.1952-87.7824a32 32 0 0 1-5.9648-44.8512z" fill="#FB553C" p-id="23544"></path><path d="M723.9168 261.7088a32 32 0 0 1 44.8512-5.9648c74.8288 57.216 112.3328 140.9024 112.3328 247.936 0 107.008-37.504 190.72-112.3328 247.936a32 32 0 1 1-38.8864-50.8416c58.24-44.544 87.2192-109.2096 87.2192-197.12 0-87.8848-28.9792-152.5504-87.2192-197.0944a32 32 0 0 1-5.9648-44.8512z" fill="#FB553C" p-id="23545"></path></svg>
                  <span style="font-size: 25px; font-weight: bold; margin-left: 10px">公告信息</span>
                </div>
                <p style="text-align: center; margin-top: 40px">欢迎使用“碳链记”！</p>
              </el-card>
              <el-card style="width: 100%; height: 300px; border-radius: 10px; margin-top: 30px">
                <div style="display: flex; align-items: center">
                  <svg t="1711772900173" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="10481" width="40" height="40"><path d="M957.6 655.2c-3.2-11.2-11.2-20-22.4-23.2l-8-1.6c-19.2-5.6-35.2-18.4-44.8-35.2-10.4-16.8-12.8-37.6-8-56.8l2.4-7.2c3.2-11.2-0.8-23.2-8.8-31.2-8.8-6.4-18.4-12.8-28-18.4-9.6-5.6-19.2-10.4-29.6-15.2-11.2-3.2-23.2-0.8-31.2 8l-7.2 5.6c-30.4 28.8-77.6 28.8-107.2 0l-5.6-5.6c-8-8-20.8-11.2-32-8-10.4 4-20 9.6-29.6 15.2-9.6 5.6-19.2 11.2-28 18.4-8.8 8-12 20-8.8 31.2l2.4 7.2c9.6 40-14.4 80.8-53.6 92.8l-7.2 1.6c-11.2 3.2-20 12-22.4 23.2-2.4 22.4-2.4 44 0 66.4 2.4 11.2 11.2 20 22.4 23.2l7.2 1.6c19.2 5.6 35.2 18.4 45.6 36 10.4 16.8 12.8 37.6 8 56.8l-2.4 7.2c-3.2 11.2 0.8 23.2 8.8 31.2 8.8 6.4 18.4 12.8 28 18.4 9.6 5.6 19.2 11.2 29.6 15.2 11.2 3.2 23.2 0 32-8l5.6-5.6c30.4-28.8 77.6-28.8 108 0l5.6 5.6c8 8 20.8 11.2 32 8 10.4-4 20-9.6 30.4-15.2 9.6-5.6 19.2-11.2 28-18.4 8.8-8 12-20 8.8-31.2l-2.4-7.2c-9.6-40.8 14.4-80.8 54.4-92l7.2-1.6c11.2-3.2 20-12 22.4-23.2 1.6-11.2 2.4-22.4 1.6-32.8-0.8-12.8-1.6-24-3.2-35.2M718.4 786.4c-52.8 0-96-43.2-96-96s43.2-96 96-96 96 43.2 96 96-43.2 96-96 96" fill="#808080" p-id="10482"></path><path d="M832.8 215.2H508.8c-10.4 0.8-20.8-2.4-28.8-10.4 0 0-13.6-24-38.4-60.8-25.6-39.2-55.2-32.8-55.2-32.8H134.4c-37.6-1.6-68.8 28-70.4 64.8v613.6c0 75.2 56.8 65.6 56.8 65.6h340.8c-90.4-142.4-48.8-331.2 93.6-421.6 102.4-65.6 234.4-63.2 335.2 4.8V272.8c3.2-28-16.8-54.4-45.6-57.6h-12M341.6 594.4H184.8c-23.2 0-42.4-19.2-42.4-42.4s19.2-42.4 42.4-42.4h156.8c23.2 0 42.4 19.2 42.4 42.4s-18.4 42.4-42.4 42.4m108.8-188H185.6c-23.2 0-42.4-19.2-42.4-42.4s19.2-42.4 42.4-42.4h264.8c23.2 0 42.4 19.2 42.4 42.4s-19.2 42.4-42.4 42.4" fill="#808080" p-id="10483"></path></svg>
                  <span style="font-size: 25px; font-weight: bold; margin-left: 10px">服务状态</span>
                </div>
              </el-card>
              <el-card style="width: 100%; height: 180px; border-radius: 10px; margin-top: 30px">
                <div style="display: flex; align-items: center">
                  <svg t="1711772109174" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="9340" width="40" height="40"><path d="M578.236052 421.891749 337.761142 421.891749c-7.384178 0-13.360286-5.982248-13.360286-13.359262 0-7.378038 5.975085-13.360286 13.360286-13.360286l240.47491 0c7.384178 0 13.360286 5.982248 13.360286 13.360286C591.596338 415.910525 585.620229 421.891749 578.236052 421.891749z" fill="#FFA500" p-id="9341"></path><path d="M685.114245 421.891749l-40.079834 0c-7.384178 0-13.360286-5.982248-13.360286-13.359262 0-7.378038 5.976108-13.360286 13.360286-13.360286l40.079834 0c7.384178 0 13.359262 5.982248 13.359262 13.360286C698.473507 415.910525 692.498422 421.891749 685.114245 421.891749z" fill="#FFA500" p-id="9342"></path><path d="M498.077408 555.490513l-146.957003 0c-7.384178 0-13.360286-5.976108-13.360286-13.361309 0-7.377014 5.975085-13.359262 13.360286-13.359262l146.957003 0c7.384178 0 13.359262 5.982248 13.359262 13.359262C511.437693 549.514405 505.462608 555.490513 498.077408 555.490513z" fill="#FFA500" p-id="9343"></path><path d="M578.236052 555.490513l-26.719548 0c-7.384178 0-13.359262-5.976108-13.359262-13.361309 0-7.377014 5.975085-13.359262 13.359262-13.359262l26.719548 0c7.384178 0 13.360286 5.982248 13.360286 13.359262C591.596338 549.514405 585.620229 555.490513 578.236052 555.490513z" fill="#FFA500" p-id="9344"></path><path d="M274.34174 729.322607c-92.161003-65.455781-143.369292-159.352311-140.472312-257.605056 5.101181-172.965353 177.446411-319.094502 376.342344-319.094502 70.425979 0 137.784086 18.415418 195.581444 49.590442-0.404206-4.168949-0.639566-8.389064-0.639566-12.662391 0-15.721053 2.975774-30.711466 8.18031-44.61922-60.888765-29.074176-130.231063-45.747927-203.122187-45.747927-227.024573 0-423.831938 169.880085-429.754834 370.953613-3.432169 116.428709 55.95643 226.77591 162.938999 302.760488 12.042267 8.55791 28.702716 5.714142 37.261649-6.315846C289.20117 754.555292 286.370705 737.868237 274.34174 729.322607z" fill="#FFA500" p-id="9345"></path><path d="M836.168566 317.011051c31.794124 46.223764 50.359968 98.958826 50.359968 153.917532 0 182.882214-168.81073 331.664795-376.316761 331.664795-4.488221 0-8.897647 1.121544-12.837376 3.273557l-121.360022 66.447365 0.678452-1.616824c10.476609-22.074761 24.684192-62.482053 2.12643-85.064374-10.437723-10.4367-27.359115-10.425444-37.795814 0-9.028631 9.027607-10.254551 22.908756-3.653203 33.255405-0.352017 6.066159-4.383844 19.178805-9.471722 30.045293l-30.528294 73.100902c-4.370541 10.476609-1.657756 22.571065 6.783497 30.163997 5.02341 4.527107 11.428284 6.848988 17.874089 6.848988 4.397147 0 8.806573-1.069355 12.837376-3.273557l182.275393-99.793844c233.781464-3.340071 422.826028-174.771489 422.826028-385.05068 0-60.086493-17.612123-117.347638-48.350195-168.190607C874.941638 311.419705 856.128153 316.482001 836.168566 317.011051z" fill="#FFA500" p-id="9346"></path><path d="M955.623191 177.502693c0 61.788252-50.100048 111.874997-111.87602 111.874997-61.788252 0-111.873973-50.085722-111.873973-111.874997 0-61.782112 50.085722-111.867834 111.873973-111.867834C905.523143 65.63486 955.623191 115.720581 955.623191 177.502693z" fill="#FFA500" p-id="9347"></path></svg>
                  <span style="font-size: 25px; font-weight: bold; margin-left: 10px">消息提醒</span>
                </div>
                <el-card style="width: 100%; border-radius: 5px; margin-top: 28px">
                  <div style="display: flex; justify-content: space-between">
                    <span>您有一条新的用户申请</span>
                    <span>3.30</span>
                  </div>
                </el-card>
              </el-card>
            </el-drawer>
          </div>
          <router-view v-else></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import avatar from "../assets/user.png"
import avatar2 from "../assets/user1.png"

import {computed, getCurrentInstance, onMounted, ref} from "vue";
import {useRouter} from "vue-router";

import * as echarts from 'echarts'

// const handleOpen = (key: string, keyPath: string[]) => {
//   console.log(key, keyPath)
// }
// const handleClose = (key: string, keyPath: string[]) => {
//   console.log(key, keyPath)
// }

const router = useRouter();
const currentInstance = getCurrentInstance()
const { proxy } = currentInstance

const UserID = sessionStorage.getItem("UserID")
const avatarUrl = ref(avatar)
const user = ref(avatar2)

const company1 = ref([])
const company2 = ref([])
const company3 = ref([])
const transaction = ref([])
const carbonAccounting = ref([])
const chartRef = ref(null)
const drawer = ref(false)

const currentRouteIsHome = computed(() => {
  return router.currentRoute.value.path === '/home'
})

const redirectToTargetPage = () => {
  window.location.href = 'http://47.97.176.174:9996/Chain01/home'
}

const handleSelect = (index) => {
  if (index === '/chainManagement'){
    // 点击了链上管理导航栏，使用路由跳转到外部链接页面
    const externalUrl = 'http://47.97.176.174:9996/Chain01/home'
    const encodedUrl = encodeURIComponent(externalUrl)
    router.push(`/chainManagement/${encodedUrl}`)
  } else if (index === '/contractWriting'){
    // 点击了合约撰写导航栏，使用路由跳转到外部链接页面
    const externalUrl = 'http://47.97.176.174:8085/'
    const encodedUrl = encodeURIComponent(externalUrl)
    router.push(`/contractWriting/${encodedUrl}`)
  }
}

const chain = () => {
  // 点击了链上管理导航栏，使用路由跳转到外部链接页面
  // const externalUrl = 'http://47.97.176.174:9996/Chain01/home'
  // const encodedUrl = encodeURIComponent(externalUrl)
  // router.push(`/chainManagement/${encodedUrl}`)
  window.location.href = 'http://47.97.176.174:9996/Chain01/home'
}

onMounted(()=>{
  proxy.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=SearchAllTransactionData").then(res=>{
    console.log(res)
    transaction.value = res.data.data
  })
  proxy.$axios.post('http://47.97.176.174:8087/chainmaker?cmb=GetAllTransactionData', {}).then(res=>{
    console.log(res)
    carbonAccounting.value = res.data.data
  })
  proxy.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=GetCompanyList").then(res=>{
    console.log(res)
    for (let i = 0; i < res.data.data.length; i++){
      if (res.data.data[i].Type == 0){
        company1.value++
      } else if (res.data.data[i].Type == 1){
        company2.value++
      } else if (res.data.data[i].Type == 2){
        company3.value++
      }
    }
  }).catch(err=>{
    console.log(err)
  })

  // reloadPage()
  // createPie()
})

const reloadPage = () => {
  // window.location.reload();
  const myChart = echarts.init(chartRef.value)
  const option = {
    xAxis: {
      type: 'category',
      data: ['2024-03-12', '2024-03-13', '2024-03-14', '2024-03-15','2024-03-16','2024-03-17','2024-03-18','2024-03-19']
    },
    yAxis: {
      type: 'value',
      name: '每1碳币可兑换的碳额度',
      axisLabel: {
        formatter: '{value}'
      }
    },
    series: [
      {
        data: ['0.0123', '0.0114', '0.0214', '0.0127','0.0134','0.0125','0.0256','0.0114'],
        type: 'line',
        label: { show: false },
        smooth: true,
        areaStyle:{},
        name: '碳额度'
      }
    ],
    tooltip: {
      trigger: 'axis'
    },
    legend: {
      data: ["碳额度"]
    },
    grid: {
      left: '7%',
      right: '1%',
      bottom: '1%',
      containLabel: true
    },
    toolbox: {
      show: false,
      feature: {
        dataZoom: {
          yAxisIndex: 'none'
        },
        dataView: {
          readOnly: false
        },
        magicType: {
          type: ['line']
        },
        restore: {},
        saveAsImage: {}
      }
    },
    color: ['#32CD32']
  }
  myChart.setOption(option)
}

const createPie = () => {
  const myChart = echarts.init(document.getElementById('pie'))
  const option = {
    tooltip: {
      trigger: 'item'
    },
    legend: {
      orient: 'horizontal',
      left: 'left',
      icon: 'rect',
      type: 'scroll'
    },
    series: [
      {
        name: '企业数',
        type: 'pie',
        radius: '50%',
        data: [
          { value: 1048, name: '电网企业' },
          { value: 735, name: '化工企业' },
          { value: 580, name: '发电企业' },
          { value: 484, name: '钢铁生产企业' },
          { value: 300, name: '电解铝生产企业' },
          { value: 248, name: '镁冶炼企业' },
          { value: 235, name: '平板玻璃生产企业' },
          { value: 180, name: '水泥生产企业' },
          { value: 84, name: '陶瓷生产企业' },
          { value: 30, name: '民航企业' },
          { value: 0, name: '其他' },
        ],
        emphasis: {
          focus: 'self',
          itemStyle: {
            shadowBlur: 10,
            shadowOffsetX: 0,
            shadowColor: 'rgba(0, 0, 0, 0.5)',
          }
        }
      }
    ]
  }
  myChart.setOption(option)
}

const typeTagType = (value) => {
  switch (value) {
    case "0":
      return 'primary'
    case "1":
      return 'primary'
    default:
      return ''
  }
}

const typeTagText = (value) => {
  switch (value) {
    case "0":
      return '电网企业'
    case "1":
      return '化工企业'
    default:
      return ''
  }
}
</script>

<style scoped>
.common-layout {
  height: 100vh;
  width: 100vw;
  position: fixed;
  background-color: #F5F5F5;
  /*background-image: url('@/assets/background.png');
  background-size: cover;
  background-repeat: no-repeat;*/
}

/* 添加带有透明度的背景颜色 */
.common-layout::before {
  content: "";
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(255, 255, 255, 0.5) /* 设置背景颜色并调整透明度 */
}

.el-menu-vertical-demo {
  border: 0;
}

.aside {
  position: fixed;
  top: 0;
  left: 0;
  height: 100vh;
  width: 275px;
}

.header {
  position: fixed;
  top: 0;
  right: 0;
  width: calc(100vw - 275px);
  padding: 0;
  display: flex;
  justify-content: right;
  align-items: center;
}

.main1 {
  position: fixed;
  bottom: 0;
  right: 0;
  height: calc(100vh - 60px);
  width: calc(100vw - 275px);
  padding: 0;
}

.aside-card {
  height: 100vh;
}

.aside :deep(.el-card__body) {
  padding: 0;
}

.aside-card-logo {
  height: 60px;
  background-color: #FCFCFC;
  display: flex;
  justify-content: center;
  align-items: center;
}

.aside-card-logo img {
  width: 28px;
}

.aside-card-logo .el-text {
  font-size: large;
  font-weight: bold;
  color: #409EFF;
}

.aside-card .el-scrollbar {
  height: calc(100vh - 80px);
}

.header-card {
  width: 1150px;
  height: 80px;
  margin: 0 0 0 0;
  padding: 0 0 5px 0;
  display: flex;
  justify-content: center;
  align-items: center;
}

.header :deep(.el-card__body) {
  padding: 0;
}

.header-card-button .el-button {
  height: 50px;
  font-size: large;
  padding: 10px;
  border: 0;
}

.header-card-button .el-button .el-icon {
  color: #000000;
  transition: all .3s ease;
  background-color: transparent;
}

.header-card-button .el-button .el-icon:hover {
  transform: scale(1.25);
}

:deep(.el-card) {
  border: 0;
  border-radius: 0;
}

.style{
  background: #fff;
  border: 1px solid #e8e8e8;
  border-radius: 5px;
  text-align: center;
  padding: 10px;
  display: inline-flex;
  justify-content: center;
  align-items: center;
}

.menu-link {
  color: black;
  text-decoration: none;
  cursor: pointer;
}

.toolbar {
  display: flex;
  align-items: center;
  justify-content: center;
}

.mybutton {
  background-color: #ffffff; /* 设置按钮背景颜色为白色 */
  color: #B22222; /* 设置按钮文字颜色 */
  border: 1px solid #B22222; /* 设置按钮边框颜色，边框宽度为1px */
}

/* 当按钮被悬停时，改变背景颜色 */
.mybutton:hover {
  background-color: #FFEFD5; /* 设置悬停时的背景颜色 */
}

 .common-layout {
   height: 100vh;
   width: 100vw;
   background-color: #F5F5F5;
 }

.aside {
  position: fixed;
  top: 0;
  left: 0;
  height: 100vh;
  width: 275px;
}

.aside-card {
  height: 100vh;
  background-color: #283342;
}

.aside-card :deep(.el-card__body) {
  padding: 0;
}

.aside-card-logo {
  height: 60px;
  background-color: #283342;
  display: flex;
  justify-content: center;
  align-items: center;
}

.aside-card-logo img {
  width: 100%;
}

.aside-card .el-scrollbar {
  height: calc(100vh - 60px);
}

.el-menu-vertical-demo span {
  font-weight: bold;
}

.el-menu-item {
  border-radius: 10px;
  margin: 0 10px;
}

.el-menu-item.is-active {
  background-color: #009099;
  border-radius: 10px;
}

:deep(.el-sub-menu__title) {
  border-radius: 10px;
  margin: 0 10px;
}

.header {
  position: fixed;
  top: 0;
  right: 0;
  width: calc(100vw - 275px);
  padding: 0;
  display: flex;
  justify-content: right;
  align-items: center;
}

.header-card {
  height: 50px;
  margin: 0 15px 0 0;
  display: flex;
  justify-content: center;
  align-items: center;
}

.header :deep(.el-card__body) {
  padding: 0;
}

.header-card-button .el-button {
  height: 50px;
  font-size: large;
  padding: 10px;
  border: 0;
}

.header-card-button .el-button .el-icon {
  color: #000000;
  transition: all .3s ease;
}

.header-card-button .el-button .el-icon:hover {
  transform: scale(1.25);
}

.header-card-button-avatar img {
  height: 25px;
  width: 25px;
  border-radius: 50%;
  margin-right: 10px;
}

.header-card-button-avatar span {
  font-size: 14px;
}

.main {
  position: fixed;
  bottom: 0;
  right: 0;
  height: calc(100vh - 60px);
  width: calc(100vw - 275px);
  padding: 0;
}

:deep(.el-card) {
  border: 0;
  border-radius: 0;
}
</style>
