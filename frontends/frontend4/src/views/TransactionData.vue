<template>
  <div style="display: flex; flex-direction: column;">
    <div style="display: flex; flex-direction: row; justify-content: space-around; align-items: center; margin-top: 30px">
      <el-card shadow="never" style="border-radius: 5px; width: 21%; background: #ecf5ff">
        <div style="display: flex; flex-direction: row">
          <div style="width: 45px; height: 45px; background: white; border: 3px solid #6EB5FF; border-radius: 50%; display: flex; justify-content: center; align-items: center">
            <svg t="1711790484715" class="icon" viewBox="0 0 1029 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4866" width="35" height="35"><path d="M704.964824 498.106533c16.466332-23.670352 26.243216-52.486432 26.243216-83.875377 0-75.642211-56.603015-137.390955-126.070352-137.390955h-20.068341V171.352764c0-6.174874-4.631156-10.80603-10.80603-10.80603h-37.049247c-6.174874 0-10.80603 4.631156-10.80603 10.80603v105.487437H467.232161V171.352764c0-6.174874-4.631156-10.80603-10.80603-10.80603h-36.020101c-6.174874 0-10.80603 4.631156-10.80603 10.80603v105.487437H288.675377c-6.174874 0-10.80603 4.631156-10.80603 10.80603v38.078392c0 5.660302 4.631156 10.80603 9.776884 10.80603 41.680402 3.60201 63.292462 31.388945 63.292462 64.836181v226.41206c0 30.874372-15.437186 55.059296-50.42814 57.632161-4.631156 1.029146-9.262312 3.60201-9.776885 9.262312l-7.20402 37.049246c-1.029146 6.174874 3.60201 12.864322 10.80603 12.864322h114.749749v105.487437c0 6.174874 4.631156 10.80603 10.80603 10.80603h36.020101c6.174874 0 10.80603-4.631156 10.80603-10.80603v-104.458292h58.661306v105.487438c0 6.174874 4.631156 10.80603 10.806031 10.80603h36.0201c6.174874 0 10.80603-4.631156 10.80603-10.80603v-105.487438h49.398995c69.467337 0 126.070352-60.205025 126.070352-135.332663 2.058291-46.311558-20.582915-87.99196-53.515578-112.176884zM422.464322 347.336683h166.207035c38.078392 0 67.923618 30.874372 67.923618 67.923619 0 38.078392-30.874372 67.923618-67.923618 67.923618H422.464322V347.336683z m199.139698 332.928644H422.978894v-135.332664h198.625126c38.078392 0 67.923618 30.874372 67.923618 67.923618 0 37.049246-30.874372 67.409045-67.923618 67.409046z" fill="#6EB5FF" p-id="4867"></path></svg>
          </div>
          <div style="display: flex; flex-direction: column; justify-content: space-between; margin-left: 10px">
            <span style="font-size: 18px; font-weight: bold">碳币交易数</span>
            <span style="font-size: 15px">Coin</span>
          </div>
        </div>
        <div style="display: flex; justify-content: space-between">
          <div style="display: flex; flex-direction: column; justify-content: center">
            <span style="font-size: 17px; font-weight: bold; margin-top: 25px">{{ coinsum }}</span>
            <p style="color: #909399">+13.4%</p>
          </div>
          <div style="display: flex; justify-content: center; align-items: center; margin-right: 15px">
            <!--          <svg t="1711791466432" class="icon" viewBox="0 0 1165 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="10902" width="100" height="100"><path d="M134.157474 418.526316h168.421052v431.157895h-168.421052v-431.157895z m747.789473-87.578948h168.421053v518.736843h-168.421053v-518.736843z m-872.421052 613.052632h1145.263158v77.473684h-1145.263158v-77.473684z m623.157894-727.578947h168.421053v633.263158h-168.421053v-633.263158z m-249.263157-215.578948h168.421052v848.842106h-168.421052v-848.842106z" fill="#67c23a" p-id="10903"></path><path d="M1.226105 501.962105l258.661053-258.661052 52.399158 52.399158L53.625263 554.368 1.226105 501.962105z m830.854737-42.186105l276.277895-276.291368 52.399158 52.399157-276.277895 276.291369-52.405895-52.399158z" fill="#67c23a" p-id="10904"></path><path d="M238.376421 315.371789l21.746526-70.844631 648.55579 199.107368-21.76 70.844632z" fill="#67c23a" p-id="10905"></path></svg>-->
            <!--        -->
            <!--          <svg t="1711951758901" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5007" width="100" height="100"><path d="M640 501.44a32 32 0 0 1-22.72-9.6L384 258.56l-233.28 233.28a32 32 0 0 1-45.44 0 32 32 0 0 1 0-45.12l256-256a32 32 0 0 1 45.44 0L640 424 873.28 192a32 32 0 0 1 45.44 0 32 32 0 0 1 0 45.12l-256 256a32 32 0 0 1-22.72 8.32z" fill="#4A8BFE" p-id="5008"></path><path d="M640 842.56a32 32 0 0 1-22.72-9.28L384 600 150.72 832a32 32 0 0 1-45.44 0 32 32 0 0 1 0-45.12l256-256a32 32 0 0 1 45.44 0L640 765.44l233.28-233.28a32 32 0 0 1 45.44 0 32 32 0 0 1 0 45.12l-256 256a32 32 0 0 1-22.72 9.28z" fill="#3BD5B3" p-id="5009"></path></svg>-->
            <!--        -->
            <svg t="1711952662368" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="6183" width="85" height="85"><path d="M461.42 704.01c-3.22 0-6.49-0.37-9.71-1.1-16.57-3.81-29.38-16.94-32.77-33.59l-69.15-338.77-49.65 145.87c-5.98 17.56-22.47 29.38-41.04 29.38H107.35C83.41 505.8 64 486.38 64 462.45s19.42-43.35 43.35-43.35h120.69l91.17-267.91c6.32-18.56 24.27-30.74 43.78-29.3 19.53 1.24 35.82 15.41 39.74 34.59l82.11 402.25 93.88-122.64c9.74-12.67 25.66-18.94 41.46-16.43 15.78 2.6 28.88 13.69 34.01 28.82l40.59 119.25h221.85c23.94 0 43.35 19.42 43.35 43.35s-19.42 43.35-43.35 43.35h-252.9c-18.57 0-35.06-11.83-41.04-29.38l-24.44-71.83-102.4 133.79a43.392 43.392 0 0 1-34.43 17zM916.65 902.22h-809.3C83.42 902.22 64 882.8 64 858.86s19.42-43.35 43.35-43.35h809.29c23.94 0 43.35 19.42 43.35 43.35s-19.41 43.36-43.34 43.36z" fill="#6EB5FF" p-id="6184"></path></svg>
          </div>
        </div>
      </el-card>
      <el-card shadow="never" style="border-radius: 5px; width: 21%; background: rgba(96,255,67,0.11)">
        <div style="display: flex; flex-direction: row">
          <div style="width: 45px; height: 45px; background: white; border: 3px solid rgba(10,183,32,0.55); border-radius: 50%; display: flex; justify-content: center; align-items: center">
            <svg t="1711792033747" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="12023" width="35" height="35"><path d="M482.148 382.087c-24.256 0-45.533-20.089-45.533-42.987 0-13.404 6.222-25.799 16.299-33.806 2.005-26.382 21.923-47.722 47.331-51.335 1.792-1.671 3.862-3.015 6.215-4.039-3.384-6.414-5.262-14.926-5.262-24.974 0-20.921 14.358-34.226 24.157-42.51-0.135-0.79-0.256-1.586-0.391-2.354-0.904-5.504-1.835-11.2-1.835-16.662 0-52.778 42.965-97.372 93.824-97.372 13.468 0 24.718 4.032 34.638 7.595 6.336 2.268 12.324 4.423 16.84 4.572 1.557-0.47 4.828-1.906 7.267-2.972 8.846-3.883 20.963-9.195 34.432-9.195 40.078 0 70.023 40.583 70.023 76.864 0 27.627-17.173 53.22-41.074 63.886-0.562 17.728-12.857 33.721-29.554 41.102 0.121 1.565 0.171 3.136 0.171 4.722 0 35.542-29.02 64.455-64.69 64.455a48.09 48.09 0 0 1-7.602-0.625c-9.109 7.025-20.928 10.787-34.41 10.787-3.549 0-6.699-0.782-9.437-1.891-11.342 12.686-26.638 20.408-42.218 20.686-0.406 0.114-0.804 0.256-1.202 0.384-1.557 0.498-3.598 1.159-6.009 1.586-0.626 0.597-1.33 1.28-1.92 1.863-5.746 5.639-15.104 14.84-28.636 16-2.198 1.578-3.933 3.164-5.74 4.821-5.247 4.807-12.444 11.4-25.684 11.4z m38.293-97.92c-3.79 3.214-7.979 4.807-12.551 4.807-10.44 0-19.591 9.678-19.591 20.708 0 0.995 0.05 2.31 0.185 3.278a17.784 17.784 0 0 1-11.577 19.122c-2.354 0.853-4.736 3.534-4.736 7.011 0 2.944 4.736 7.069 9.372 7.41 0.654-0.547 1.572-1.387 2.297-2.055 3.463-3.164 8.2-7.51 15.836-11.84a17.832 17.832 0 0 1 10.532-2.226c1.664-0.54 5.191-4.01 7.104-5.888 5.106-5.013 11.42-11.221 20.914-11.598 0.334-0.1 0.675-0.213 1.01-0.327 2.588-0.832 6.5-2.09 11.235-2.09 9.23 0 17.55-8.207 21.703-16.342a17.764 17.764 0 0 1 15.83-9.693c7.95 0 13.596 4.075 17.336 6.777 0.157 0.121 0.32 0.235 0.477 0.356 5.753-0.455 9.842-2.368 12.188-5.696a17.757 17.757 0 0 1 19.556-6.813 17.477 17.477 0 0 1 4.672 2.148c1.024 0.192 1.963 0.313 2.78 0.313 16.064 0 29.135-12.964 29.135-28.9 0-3.932-0.74-7.594-2.262-11.185a17.709 17.709 0 0 1-1.358-8.32c0.711-9.693 9.195-17.074 18.71-16.463 7.103-0.042 14.321-5.738 14.321-11.114 0-2.539 0.135-5.028 0.27-7.41 0.071-1.579 0.171-3.08 0.171-4.501 0-8.64 6.215-16.036 14.734-17.515 12.494-2.176 25.87-16.043 25.87-33.223 0-20.075-17.713-41.309-34.467-41.309-6.016 0-13.518 3.293-20.146 6.201-7.011 3.08-13.639 5.988-20.679 5.988-11.115 0-20.572-3.392-29.717-6.67-7.9-2.83-15.36-5.504-22.642-5.504-31.04 0-58.269 28.892-58.269 61.816 0 2.568 0.733 6.998 1.366 10.909 0.853 5.163 1.728 10.503 1.728 15.417 0 4.729-1.885 9.294-5.234 12.63-1.742 1.734-3.996 3.654-6.514 5.766-5.916 4.957-13.276 11.122-13.276 16.804 0 2.41 0.291 4.622 0.618 6.222a17.816 17.816 0 0 1 6.55 18.368c-3.933 16.156-8.868 27.392-20.274 32.078a17.815 17.815 0 0 1-3.207 2.553zM849.188 945.778H172.8a17.773 17.773 0 0 1-16.427-10.994 17.8 17.8 0 0 1 3.89-19.392c1.949-1.941 195.172-197.803 195.172-485.177 0-9.82 7.957-17.778 17.777-17.778h275.556c9.82 0 17.778 7.958 17.778 17.778 0 287.936 193.223 483.229 195.171 485.17a17.776 17.776 0 0 1 3.911 19.392 17.791 17.791 0 0 1-16.44 11z m-636.601-35.556h596.821C754.709 844.487 637.078 677.148 631.218 448H390.763c-5.86 229.148-123.485 396.487-178.176 462.222z" fill="rgba(10,183,32,0.55)" p-id="12024"></path></svg>
          </div>
          <div style="display: flex; flex-direction: column; justify-content: space-between; margin-left: 10px">
            <span style="font-size: 18px; font-weight: bold">碳排放交易数</span>
            <span style="font-size: 15px">Emission</span>
          </div>
        </div>
        <div style="display: flex; justify-content: space-between">
          <div style="display: flex; flex-direction: column; justify-content: center">
            <span style="font-size: 17px; font-weight: bold; margin-top: 25px">{{ creditsum }}</span>
            <p style="color: #909399">-22.4%</p>
          </div>
          <div style="display: flex; justify-content: center; align-items: center">
            <!--          <svg t="1711791466432" class="icon" viewBox="0 0 1165 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="10902" width="100" height="100"><path d="M134.157474 418.526316h168.421052v431.157895h-168.421052v-431.157895z m747.789473-87.578948h168.421053v518.736843h-168.421053v-518.736843z m-872.421052 613.052632h1145.263158v77.473684h-1145.263158v-77.473684z m623.157894-727.578947h168.421053v633.263158h-168.421053v-633.263158z m-249.263157-215.578948h168.421052v848.842106h-168.421052v-848.842106z" fill="#EE6363" p-id="10903"></path><path d="M1.226105 501.962105l258.661053-258.661052 52.399158 52.399158L53.625263 554.368 1.226105 501.962105z m830.854737-42.186105l276.277895-276.291368 52.399158 52.399157-276.277895 276.291369-52.405895-52.399158z" fill="#EE6363" p-id="10904"></path><path d="M238.376421 315.371789l21.746526-70.844631 648.55579 199.107368-21.76 70.844632z" fill="#EE6363" p-id="10905"></path></svg>-->
            <!--        -->
            <!--          <svg t="1711951758901" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5007" width="100" height="100"><path d="M640 501.44a32 32 0 0 1-22.72-9.6L384 258.56l-233.28 233.28a32 32 0 0 1-45.44 0 32 32 0 0 1 0-45.12l256-256a32 32 0 0 1 45.44 0L640 424 873.28 192a32 32 0 0 1 45.44 0 32 32 0 0 1 0 45.12l-256 256a32 32 0 0 1-22.72 8.32z" fill="#4A8BFE" p-id="5008"></path><path d="M640 842.56a32 32 0 0 1-22.72-9.28L384 600 150.72 832a32 32 0 0 1-45.44 0 32 32 0 0 1 0-45.12l256-256a32 32 0 0 1 45.44 0L640 765.44l233.28-233.28a32 32 0 0 1 45.44 0 32 32 0 0 1 0 45.12l-256 256a32 32 0 0 1-22.72 9.28z" fill="#3BD5B3" p-id="5009"></path></svg>-->
            <!--        -->
            <svg t="1711953216961" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="17663" width="85" height="85"><path d="M534.840892 721.739223v285.227143a16.574788 16.574788 0 0 0 4.92764 12.03112 16.894765 16.894765 0 0 0 11.967125 4.927639h124.98286a16.574788 16.574788 0 0 0 11.967125-4.927639 17.022755 17.022755 0 0 0 4.92764-12.03112v-417.249489L566.518576 699.468851l-31.997661-27.262006zM0.927934 1006.966366a16.574788 16.574788 0 0 0 4.92764 12.03112 16.894765 16.894765 0 0 0 11.967125 4.927639h124.982861a16.574788 16.574788 0 0 0 11.967125-4.927639 17.022755 17.022755 0 0 0 4.927639-12.03112V679.630302L0.927934 811.140685z m266.988477-416.865517v416.737526a16.702779 16.702779 0 0 0 4.92764 12.03112 16.894765 16.894765 0 0 0 11.967125 4.92764H409.60205a16.574788 16.574788 0 0 0 11.967125-4.92764 17.022755 17.022755 0 0 0 4.92764-12.03112V592.212694l-79.418193-67.835039-79.226206 65.723194z m534.040948 416.865517a16.574788 16.574788 0 0 0 4.92764 12.03112 16.894765 16.894765 0 0 0 11.967125 4.927639h124.98286a16.574788 16.574788 0 0 0 11.967125-4.927639 17.022755 17.022755 0 0 0 4.92764-12.03112V383.971922l-158.77239 137.461948zM744.297576 0a5.439602 5.439602 0 0 0-3.83972 9.279321l114.8716 113.335713-310.633285 268.780345-210.800585-180.338813L0.032 487.324364v119.415268l333.799591-276.395788L544.632176 510.554666l375.140568-324.200293 94.841065 93.625153a5.439602 5.439602 0 0 0 9.279321-3.839719V5.567593a5.503598 5.503598 0 0 0-5.439602-5.439602z" fill="rgba(10,183,32,0.55)" p-id="17664"></path></svg>
          </div>
        </div>
      </el-card>
      <el-card shadow="never" style="border-radius: 5px; width: 21%; background: rgba(255,229,57,0.18)">
        <div style="display: flex; flex-direction: row">
          <div style="width: 45px; height: 45px; background: white; border: 3px solid #FFA500; border-radius: 50%; display: flex; justify-content: center; align-items: center">
            <svg t="1711792314350" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="14236" width="32" height="32"><path d="M512 5.8368a505.9584 505.9584 0 0 1 501.1456 436.736H720.6912L535.6544 257.7408l-2.4576-2.2528a34.5088 34.5088 0 0 0-46.2848 2.2528l-113.7664 113.664-2.2528 2.4576a34.5088 34.5088 0 0 0 2.2528 45.9776l2.4576 2.2528a34.5088 34.5088 0 0 0 46.3872-2.2528l89.2928-89.2928L681.984 501.76l2.56 2.2528A34.5088 34.5088 0 0 0 706.56 512h311.296a505.856 505.856 0 0 1-497.152 505.856H512A505.9584 505.9584 0 0 1 8.192 557.3632h283.7504l184.9344 184.9344 2.4576 2.2528a34.5088 34.5088 0 0 0 46.2848-2.2528l113.8688-113.7664 2.2528-2.4576a34.5088 34.5088 0 0 0-2.2528-46.3872l-2.4576-2.2528a34.5088 34.5088 0 0 0-46.3872 2.2528L501.76 669.0816 330.6496 498.3808l-2.9696-2.2528a34.5088 34.5088 0 0 0-21.9136-7.8848H6.7584A505.856 505.856 0 0 1 512 5.8368z m0 0" fill="#FFA500" p-id="14237"></path></svg>
          </div>
          <div style="display: flex; flex-direction: column; justify-content: space-between; margin-left: 10px">
            <span style="font-size: 18px; font-weight: bold">交易总数量</span>
            <span style="font-size: 15px">Total</span>
          </div>
        </div>
        <div style="display: flex; justify-content: space-between">
          <div style="display: flex; flex-direction: column; justify-content: center">
            <span style="font-size: 17px; font-weight: bold; margin-top: 25px">{{ sum }}</span>
            <p style="color: #909399">+12.4%</p>
          </div>
          <div style="display: flex; justify-content: center; align-items: center">
            <!--          <svg t="1711791466432" class="icon" viewBox="0 0 1165 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="10902" width="100" height="100"><path d="M134.157474 418.526316h168.421052v431.157895h-168.421052v-431.157895z m747.789473-87.578948h168.421053v518.736843h-168.421053v-518.736843z m-872.421052 613.052632h1145.263158v77.473684h-1145.263158v-77.473684z m623.157894-727.578947h168.421053v633.263158h-168.421053v-633.263158z m-249.263157-215.578948h168.421052v848.842106h-168.421052v-848.842106z" fill="#67c23a" p-id="10903"></path><path d="M1.226105 501.962105l258.661053-258.661052 52.399158 52.399158L53.625263 554.368 1.226105 501.962105z m830.854737-42.186105l276.277895-276.291368 52.399158 52.399157-276.277895 276.291369-52.405895-52.399158z" fill="#67c23a" p-id="10904"></path><path d="M238.376421 315.371789l21.746526-70.844631 648.55579 199.107368-21.76 70.844632z" fill="#67c23a" p-id="10905"></path></svg>-->
            <!--        -->
            <svg t="1711951758901" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5007" width="85" height="85"><path d="M640 501.44a32 32 0 0 1-22.72-9.6L384 258.56l-233.28 233.28a32 32 0 0 1-45.44 0 32 32 0 0 1 0-45.12l256-256a32 32 0 0 1 45.44 0L640 424 873.28 192a32 32 0 0 1 45.44 0 32 32 0 0 1 0 45.12l-256 256a32 32 0 0 1-22.72 8.32z" fill="#FFA500" p-id="5008"></path><path d="M640 842.56a32 32 0 0 1-22.72-9.28L384 600 150.72 832a32 32 0 0 1-45.44 0 32 32 0 0 1 0-45.12l256-256a32 32 0 0 1 45.44 0L640 765.44l233.28-233.28a32 32 0 0 1 45.44 0 32 32 0 0 1 0 45.12l-256 256a32 32 0 0 1-22.72 9.28z" fill="#FFA500" p-id="5009"></path></svg>
          </div>
        </div>
      </el-card>
      <el-card shadow="never" style="border-radius: 5px; width: 21%; background: rgba(255,139,225,0.13)">
        <div style="display: flex; flex-direction: row">
          <div style="width: 45px; height: 45px; background: white; border: 3px solid rgba(200,33,146,0.55); border-radius: 50%; display: flex; justify-content: center; align-items: center">
            <svg t="1711792875816" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="15336" width="32" height="32"><path d="M554.666667 422.4c0-21.333333-21.333333-42.666667-42.666667-42.666667-25.6 0-42.666667 21.333333-42.666667 42.666667V640c0 21.333333 17.066667 42.666667 42.666667 42.666667 21.333333 0 42.666667-21.333333 42.666667-42.666667V422.4z" p-id="15337" fill="rgba(200,33,146,0.55)"></path><path d="M512 768m-42.666667 0a42.666667 42.666667 0 1 0 85.333334 0 42.666667 42.666667 0 1 0-85.333334 0Z" p-id="15338" fill="rgba(200,33,146,0.55)"></path><path d="M964.266667 563.2h-17.066667l-68.266667-8.533333h-17.066666c-8.533333 4.266667-8.533333 8.533333-12.8 17.066666s0 12.8 8.533333 17.066667l12.8 12.8-68.266667 55.466667-4.266666-4.266667c-8.533333-8.533333-17.066667-17.066667-21.333334-25.6-8.533333-12.8-21.333333-12.8-34.133333-4.266667-12.8 8.533333-25.6 21.333333-38.4 29.866667L610.133333 725.333333c-4.266667 4.266667-12.8 8.533333-12.8 17.066667s4.266667 12.8 8.533334 17.066667l4.266666 4.266666c8.533333 8.533333 12.8 12.8 17.066667 12.8 4.266667 0 8.533333-4.266667 17.066667-8.533333l102.4-76.8 25.6 29.866667c12.8 12.8 21.333333 12.8 34.133333 4.266666l162.133333-123.733333 8.533334-8.533333c4.266667-4.266667 4.266667-8.533333 4.266666-17.066667-4.266667-8.533333-8.533333-12.8-17.066666-12.8zM853.333333 725.333333h85.333334v213.333334h-85.333334zM725.333333 768h85.333334v170.666667h-85.333334zM597.333333 853.333333h85.333334v85.333334h-85.333334z" p-id="15339" fill="rgba(200,33,146,0.55)"></path><path d="M512 213.333333l187.733333 341.333334h98.133334L550.4 106.666667c-17.066667-25.6-59.733333-25.6-72.533333 0l-426.666667 768c-8.533333 12.8-8.533333 29.866667 0 42.666666 4.266667 12.8 17.066667 21.333333 34.133333 21.333334h469.333334v-85.333334H157.866667L512 213.333333z" p-id="15340" fill="rgba(200,33,146,0.55)"></path></svg>
          </div>
          <div style="display: flex; flex-direction: column; justify-content: space-between; margin-left: 10px">
            <span style="font-size: 18px; font-weight: bold">交易失败数</span>
            <span style="font-size: 15px">Failure</span>
          </div>
        </div>
        <div style="display: flex; justify-content: space-between">
          <div style="display: flex; flex-direction: column; justify-content: center">
            <span style="font-size: 17px; font-weight: bold; margin-top: 25px">13</span>
            <p style="color: #909399">-7.4%</p>
          </div>
          <div style="display: flex; justify-content: center; align-items: center">
            <!--          <svg t="1711791466432" class="icon" viewBox="0 0 1165 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="10902" width="100" height="100"><path d="M134.157474 418.526316h168.421052v431.157895h-168.421052v-431.157895z m747.789473-87.578948h168.421053v518.736843h-168.421053v-518.736843z m-872.421052 613.052632h1145.263158v77.473684h-1145.263158v-77.473684z m623.157894-727.578947h168.421053v633.263158h-168.421053v-633.263158z m-249.263157-215.578948h168.421052v848.842106h-168.421052v-848.842106z" fill="#EE6363" p-id="10903"></path><path d="M1.226105 501.962105l258.661053-258.661052 52.399158 52.399158L53.625263 554.368 1.226105 501.962105z m830.854737-42.186105l276.277895-276.291368 52.399158 52.399157-276.277895 276.291369-52.405895-52.399158z" fill="#EE6363" p-id="10904"></path><path d="M238.376421 315.371789l21.746526-70.844631 648.55579 199.107368-21.76 70.844632z" fill="#EE6363" p-id="10905"></path></svg>-->
            <!--        -->
            <!--          <svg t="1711951758901" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5007" width="100" height="100"><path d="M640 501.44a32 32 0 0 1-22.72-9.6L384 258.56l-233.28 233.28a32 32 0 0 1-45.44 0 32 32 0 0 1 0-45.12l256-256a32 32 0 0 1 45.44 0L640 424 873.28 192a32 32 0 0 1 45.44 0 32 32 0 0 1 0 45.12l-256 256a32 32 0 0 1-22.72 8.32z" fill="#4A8BFE" p-id="5008"></path><path d="M640 842.56a32 32 0 0 1-22.72-9.28L384 600 150.72 832a32 32 0 0 1-45.44 0 32 32 0 0 1 0-45.12l256-256a32 32 0 0 1 45.44 0L640 765.44l233.28-233.28a32 32 0 0 1 45.44 0 32 32 0 0 1 0 45.12l-256 256a32 32 0 0 1-22.72 9.28z" fill="#3BD5B3" p-id="5009"></path></svg>-->
            <!--        -->
            <svg t="1711956877766" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="36330" width="85" height="85"><path d="M93.1 615.2c20.7 15.5 46.5 15.5 62.1 0 15.5-20.6 15.5-46.4 0-61.9s-41.4-15.5-62.1 0c-15.5 15.5-15.5 41.3 0 61.9z" fill="rgba(200,33,146,0.55)" p-id="36331"></path><path d="M356.8 367.6c20.7 15.5 46.5 15.5 62.1 0 15.5-20.6 15.5-46.4 0-61.9s-41.4-15.5-62.1 0c-15.5 15.5-15.5 41.3 0 61.9z" fill="rgba(200,33,146,0.55)" p-id="36332"></path><path d="M62.1 754.5h124.1v170.2H62.1V754.5z" fill="rgba(200,33,146,0.55)" p-id="36333"></path><path d="M325.8 506.9h124.1v417.8H325.8V506.9z" fill="rgba(200,33,146,0.55)" p-id="36334"></path><path d="M574.1 630.7h124.1v294H574.1v-294z" fill="rgba(200,33,146,0.55)" p-id="36335"></path><path d="M822.3 336.7h124.1v588H822.3v-588z" fill="rgba(200,33,146,0.55)" p-id="36336"></path><path d="M605.1 491.4c20.7 15.5 46.5 15.5 62.1 0 15.5-20.6 15.5-46.4 0-61.9s-41.4-15.5-62.1 0c-15.5 15.5-15.5 41.3 0 61.9z" fill="rgba(200,33,146,0.55)" p-id="36337"></path><path d="M853.3 197.4c20.7 15.5 46.5 15.5 62.1 0 15.5-20.6 15.5-46.4 0-61.9s-41.4-15.5-62.1 0c-15.5 15.5-15.5 41.3 0 61.9z" fill="rgba(200,33,146,0.55)" p-id="36338"></path><path d="M884.4 166.4" fill="rgba(200,33,146,0.55)" p-id="36339"></path><path d="M418.9 352.1l170.7 77.4" fill="rgba(200,33,146,0.55)" p-id="36340"></path><path d="M170.7 537.8l170.6-170.2" fill="rgba(200,33,146,0.55)" p-id="36341"></path><path d="M682.7 414l170.6-201.1" fill="rgba(200,33,146,0.55)" p-id="36342"></path></svg>
          </div>
        </div>
      </el-card>
    </div>
    <el-card shadow="never" style="margin: 30px 15px 0px 15px">
      <div style="display: flex; align-items: center; margin-top: 10px; margin-left: 20px">
        <text>公司编号：</text>
        <el-input v-model="CompanyId" placeholder="请输入公司编号" style="width: 20%; margin-left: 5px; height: 40px" />
        <el-button type="primary" icon="Search" @click="search" style="width: 13%; margin-left: 20px; font-size: 16px; padding-top: 18px; padding-bottom: 18px">搜索</el-button>
        <el-button type="success" icon="List" @click="getTableData" style="width: 180px; height: 37px; font-size: 16px; margin-left: 50px">查询所有</el-button>
        <div class="style" style="display: flex; align-items: center; margin-left: 8%; color: #696969">
          <i class="pi pi-server"></i>
          <text style="margin-left: 8px">交易总数：</text>
          <text style="font-size: 22px; font-weight: bold; margin-left: 5px">{{ sum }}</text>
        </div>
      </div>
      <el-table :data="tableData" border height="390" :header-cell-style="{ background:'#f4f4f5' }" style="width: 98%; margin-left: 10px; margin-top: 20px">
        <el-table-column label="编号" prop="Id" align="center" width="80"></el-table-column>
        <el-table-column label="公司编号" prop="CompanyId" align="center" width="150"></el-table-column>
        <el-table-column label="碳币" prop="CarbonCoin" align="center" width="120"></el-table-column>
        <el-table-column label="碳额度" prop="CarbonCredit" align="center" width="120"></el-table-column>
        <el-table-column label="交易编号" prop="Txid" align="center"></el-table-column>
        <el-table-column label="操作" align="center" width="120">
          <template #default="scope">
            <el-button type="warning" icon="InfoFilled" @click="traceSource(scope.row)">详情</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-card>
  </div>
  <el-dialog v-model="visible" title="交易溯源详情" width="700px">
    <el-descriptions border :column="2">
      <el-descriptions-item label="区块高度" width="120">{{ block_height }}</el-descriptions-item>
      <el-descriptions-item label="区块时间戳" width="100">{{ block_timestamp }}</el-descriptions-item>
      <el-descriptions-item label="区块链ID">{{ chain_id }}</el-descriptions-item>
      <el-descriptions-item label="交易发起组织">{{ org_id }}</el-descriptions-item>
      <el-descriptions-item label="合约名称">{{ contract_name }}</el-descriptions-item>
      <el-descriptions-item label="合约调用方法">{{ method }}</el-descriptions-item>
      <el-descriptions-item label="合约调用参数" :span="2">{{ parameters }}</el-descriptions-item>
      <el-descriptions-item label="合约执行结果">{{ contract_result_message }}</el-descriptions-item>
      <el-descriptions-item label="消耗GAS量">{{ gas_used }}</el-descriptions-item>
      <el-descriptions-item label="交易ID" :span="2">{{ tx_id }}</el-descriptions-item>
    </el-descriptions>
  </el-dialog>
</template>

<script setup>
import {getCurrentInstance, onMounted, ref} from "vue";
import {useRouter} from "vue-router";

const currentInstance = getCurrentInstance()
const { proxy } = currentInstance
const router = useRouter()

const tableData = ref([])
const sum = ref(0)
const coinsum = ref(0)
const creditsum = ref(0)
const CompanyId = ref('')

const visible = ref(false)
const block_height = ref(-1)
const block_timestamp = ref(-1)
const chain_id = ref('')
const contract_name = ref('')
const method = ref('')
const parameters = ref([])
const tx_id = ref('')
const contract_result_message = ref('')
const gas_used = ref(0)
const org_id = ref('')

onMounted(()=>{
  proxy.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=SearchAllTransactionData").then(res=>{
    console.log(res)
    tableData.value = res.data.data
    sum.value = res.data.data.length
    for (let i = 0; i < res.data.data.length; i++){
      coinsum.value += res.data.data[i].CarbonCoin
      creditsum.value += res.data.data[i].CarbonCredit
    }
  })
})

const search = () => {
  proxy.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=SearchAllTransactionDataByCompanyId",{
    CompanyId: CompanyId.value
  }).then(res=>{
    tableData.value = res.data.data
    console.log(res)
  })
}

const getTableData = () => {
  proxy.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=SearchAllTransactionData").then(res=>{
    console.log(res)
    tableData.value = res.data.data
    sum.value = res.data.data.length
  })
}

const traceSource = (row) => {
  visible.value = true
  proxy.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=TxIdToSearchTx",{
    Txid: row.Txid
  }).then(res=>{
    console.log(res)
    block_height.value = res.data.data.block_height
    block_timestamp.value = res.data.data.block_timestamp
    chain_id.value = res.data.data.transaction.payload.chain_id
    contract_name.value = res.data.data.transaction.payload.contract_name
    method.value = res.data.data.transaction.payload.method
    const data = []
    for (let i = 0; i < res.data.data.transaction.payload.parameters.length; i++){
      data.push(res.data.data.transaction.payload.parameters[i].key)
    }
    parameters.value = [...data]
    tx_id.value = res.data.data.transaction.payload.tx_id
    contract_result_message.value = res.data.data.transaction.result.contract_result.message
    gas_used.value = res.data.data.transaction.result.contract_result.gas_used
    org_id.value = res.data.data.transaction.sender.signer.org_id
  })
}
</script>

<style scoped>
.style{
  background: rgba(0,0,0,0.04);
  border-radius: 10px;
  padding: 15px;
}
</style>