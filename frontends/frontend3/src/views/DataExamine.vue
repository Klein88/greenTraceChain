<script setup>
import {computed, onMounted, ref, watch} from "vue";
import axios from "axios";
import {
  Refresh,
  Plus,
  Edit,
  Delete,
  Download,
  Search,
  CircleCheckFilled,
  CircleCloseFilled,
  InfoFilled,
  CaretBottom,
  CaretTop,
  CircleClose,
  CircleCheck,
  Warning
} from '@element-plus/icons-vue'
import * as XLSX from "xlsx";
import {ElMessage} from "element-plus";

const tableData = ref([])
const searchQuery = ref('')
const originalTableData = ref([])
const currentCompanyData = ref({})
const dialogVisibleForPGE = ref(false)
const dialogVisibleForCE = ref(false)
const dialogVisibleForFiles = ref(false)
const fileList = ref([])
const currentRow = ref({})
const currentPage = ref(1)
const pageSize = ref(15)
const small = ref(false)
const background = ref(true)
const disabled = ref(false)

onMounted(async () => {
  tableData.value = await axios.post('http://47.97.176.174:8087/chainmaker?cmb=GetAllTransactionData', {})
      .then(function (response) {
        console.log(response)
        return response.data.data.map(item => {
          item.createAt = item.createAt.split('T')[0];
          item.updateAt = item.updateAt.split('T')[0];
          return item;
        })
      })
      .catch(function (error) {
        console.log(error)
      })
  originalTableData.value = [...tableData.value]
})

watch(searchQuery, () => {
  if (searchQuery.value) {
    tableData.value = originalTableData.value.filter(item => item.CompanyId.includes(searchQuery.value))
  } else {
    tableData.value = originalTableData.value
  }
})

const totalCount = computed(() => {
  return tableData.value.length;
})
const unreviewedCount = computed(() => {
  return tableData.value.filter(item => item.State === 0).length;
})
const reviewedPassedCount = computed(() => {
  return tableData.value.filter(item => item.State === 1).length;
})
const reviewedNotPassedCount = computed(() => {
  return tableData.value.filter(item => item.State === 2).length;
})

const refresh = async () => {
  tableData.value = await axios.post('http://47.97.176.174:8087/chainmaker?cmb=GetAllTransactionData', {})
      .then(function (response) {
        ElMessage.success('刷新成功')
        console.log(response)
        return response.data.data.map(item => {
          item.createAt = item.createAt.split('T')[0];
          item.updateAt = item.updateAt.split('T')[0];
          return item;
        })
      })
      .catch(function (error) {
        ElMessage.error('刷新失败')
        console.log(error)
      })
  originalTableData.value = [...tableData.value];
}

const exportToExcel = () => {
  const ws = XLSX.utils.json_to_sheet(tableData.value)
  const wb = XLSX.utils.book_new()
  XLSX.utils.book_append_sheet(wb, ws, "Sheet1")
  XLSX.writeFile(wb, "TableData.xlsx")
  ElMessage.success('导出成功');
}

const stateTagType = (value) => {
  switch (value) {
    case 0:
      return 'danger'
    case 1:
      return 'success'
    case 2:
      return 'warning'
    default:
      return ''
  }
}
const stateTagText = (value) => {
  switch (value) {
    case 0:
      return '未审核'
    case 1:
      return '已审核通过'
    case 2:
      return '未审核通过'
    default:
      return ''
  }
}
const stateOptions = [
  { text: '未审核', value: 0 },
  { text: '已审核通过', value: 1 },
  { text: '未审核通过', value: 2 }
]
const filterState = (value, row) => {
  return row.State === value
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
const typeOptions = [
  { text: '电网企业', value: "0" },
  { text: '化工企业', value: "1" }
]
const filterType = (value, row) => {
  return row.Type === value
}

const showFiles = (index, row) => {
  try {
    fileList.value = JSON.parse(row.CompanyFileUrl)
  } catch (e) {
    console.error('解析CompanyFileUrl失败:', e)
  }
  currentRow.value = row
  dialogVisibleForFiles.value = true
}

const fileIcon = (fileName) => {
  const extension = fileName.split('.').pop().toLowerCase();
  switch (extension) {
    case 'jpg':
    case 'jpeg':
    case 'png':
    case 'gif':
    case 'bmp':
    case 'svg':
    case 'tif':
    case 'tiff':
    case 'ico':
    case 'webp':
    case 'psd':
      return '<svg t="1712414559675" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="10505" width="32" height="32"><path d="M788.14 928.23h-550c-77.32 0-140-62.68-140-140v-550c0-77.32 62.68-140 140-140h550c77.32 0 140 62.68 140 140v550c0 77.32-62.68 140-140 140z" fill="#FFC757" p-id="10506"></path><path d="M770.19 644.92l-67.56-117.01c-15.4-26.68-53.91-26.68-69.31 0l-64.01 110.88L450.3 432.65c-15.4-26.67-53.89-26.67-69.29 0L258.45 644.94c-15.4 26.67 3.85 60.01 34.65 60.01h442.44c30.8-0.01 50.05-33.35 34.65-60.03z" fill="#EF9F2B" p-id="10507"></path><path d="M595.103921 388.209544a55.27 55.27 0 1 0 100.661393-45.676862 55.27 55.27 0 1 0-100.661393 45.676862Z" fill="#EF9F2B" p-id="10508"></path><path d="M770.19 632.92l-67.56-117.01c-15.4-26.68-53.91-26.68-69.31 0l-64.01 110.88L450.3 420.65c-15.4-26.67-53.89-26.67-69.29 0L258.45 632.94c-15.4 26.67 3.85 60.01 34.65 60.01h442.44c30.8-0.01 50.05-33.35 34.65-60.03z" fill="#FFFFFF" p-id="10509"></path><path d="M595.104335 376.209454a55.27 55.27 0 1 0 100.661392-45.676862 55.27 55.27 0 1 0-100.661392 45.676862Z" fill="#FFFFFF" p-id="10510"></path></svg>';
    case 'doc':
    case 'docx':
      return '<svg t="1712414485474" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="9837" width="32" height="32"><path d="M658.78 64.39H269.46c-78.06 0-141.33 63.28-141.33 141.33v612.95c0 78.06 63.28 141.33 141.33 141.33h485.49c78.06 0 141.33-63.28 141.33-141.33V301.88c0-31.83-12.64-62.35-35.15-84.85l-117.5-117.5a120.008 120.008 0 0 0-84.85-35.14z" fill="#53B7F4" p-id="9838"></path><path d="M647.85 360.65c-13.25 0-24 10.75-24 24v226.26l-94.35-93.29c-9.33-9.22-24.34-9.25-33.69-0.05l-95.26 93.6V383.99c0-13.25-10.75-24-24-24s-24 10.75-24 24V668.4c0 9.67 5.8 18.39 14.72 22.13a24.002 24.002 0 0 0 26.1-5.01L512.58 568.4l118.4 117.07a23.993 23.993 0 0 0 16.88 6.93c3.11 0 6.25-0.61 9.24-1.85a24 24 0 0 0 14.75-22.15V384.65c0-13.25-10.74-24-24-24zM861.13 217.03l-117.5-117.5a120.001 120.001 0 0 0-61.75-32.9v131.88c0 41.87 33.94 75.81 75.81 75.81h135.37a119.975 119.975 0 0 0-31.93-57.29z" fill="#29A3D3" p-id="9839"></path><path d="M647.86 680.4c-6.2 0-12.3-2.4-16.88-6.93L512.58 556.4 393.37 673.52a23.99 23.99 0 0 1-26.1 5.01 23.992 23.992 0 0 1-14.72-22.13V371.99c0-13.25 10.75-24 24-24s24 10.75 24 24v227.19l95.26-93.6c9.36-9.19 24.37-9.17 33.69 0.05l94.35 93.29V372.65c0-13.25 10.75-24 24-24s24 10.75 24 24V656.4c0 9.68-5.82 18.42-14.75 22.15a24.112 24.112 0 0 1-9.24 1.85z" fill="#FFFFFF" p-id="9840"></path></svg>';
    case 'xls':
    case 'xlsx':
      return '<svg t="1712414504057" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="10003" width="32" height="32"><path d="M658.68 63.95H269.36c-78.06 0-141.33 63.28-141.33 141.33v612.95c0 78.06 63.28 141.33 141.33 141.33h485.49c78.06 0 141.33-63.28 141.33-141.33V301.44c0-31.83-12.64-62.35-35.15-84.85l-117.5-117.5a120.008 120.008 0 0 0-84.85-35.14z" fill="#53D39C" p-id="10004"></path><path d="M655.49 388.82c-9.37-9.37-24.57-9.37-33.94 0L510.8 499.57 400.05 388.82c-9.37-9.37-24.57-9.37-33.94 0-9.37 9.37-9.37 24.57 0 33.94l110.75 110.75L366.1 644.26c-9.37 9.37-9.37 24.57 0 33.94 4.69 4.69 10.83 7.03 16.97 7.03s12.28-2.34 16.97-7.03L510.8 567.45 621.54 678.2c4.69 4.69 10.83 7.03 16.97 7.03s12.28-2.34 16.97-7.03c9.37-9.37 9.37-24.57 0-33.94L544.74 533.51l110.75-110.75c9.37-9.37 9.37-24.57 0-33.94zM861.03 216.59l-117.5-117.5a120.001 120.001 0 0 0-61.75-32.9v131.88c0 41.87 33.94 75.81 75.81 75.81h135.37a120.058 120.058 0 0 0-31.93-57.29z" fill="#25BF79" p-id="10005"></path><path d="M544.74 521.51l110.75-110.75c9.37-9.37 9.37-24.57 0-33.94-9.37-9.37-24.57-9.37-33.94 0L510.8 487.57 400.05 376.82c-9.37-9.37-24.57-9.37-33.94 0-9.37 9.37-9.37 24.57 0 33.94l110.75 110.75L366.1 632.26c-9.37 9.37-9.37 24.57 0 33.94 4.69 4.69 10.83 7.03 16.97 7.03s12.28-2.34 16.97-7.03L510.8 555.45 621.54 666.2c4.69 4.69 10.83 7.03 16.97 7.03s12.28-2.34 16.97-7.03c9.37-9.37 9.37-24.57 0-33.94L544.74 521.51z" fill="#FFFFFF" p-id="10006"></path></svg>';
    case 'ppt':
    case 'pptx':
      return '<svg t="1712414513128" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="10171" width="32" height="32"><path d="M658.58 63.94H269.26c-78.06 0-141.33 63.28-141.33 141.33v612.95c0 78.06 63.28 141.33 141.33 141.33h485.49c78.06 0 141.33-63.28 141.33-141.33V301.44c0-31.83-12.64-62.35-35.15-84.85l-117.5-117.5a120.017 120.017 0 0 0-84.85-35.15z" fill="#F98950" p-id="10172"></path><path d="M860.93 216.59l-117.5-117.5a120.001 120.001 0 0 0-61.75-32.9v131.88c0 41.87 33.94 75.81 75.81 75.81h135.37a119.93 119.93 0 0 0-31.93-57.29zM556.02 348.38H380.37c-13.25 0-24 10.75-24 24v302.74c0 13.25 10.75 24 24 24s24-10.75 24-24V595.6h151.65c68.16 0 123.61-55.45 123.61-123.61s-55.45-123.61-123.61-123.61z m0 199.22H404.37V396.38h151.65c41.69 0 75.61 33.92 75.61 75.61s-33.92 75.61-75.61 75.61z" fill="#F26C38" p-id="10173"></path><path d="M556.02 336.38H380.37c-13.25 0-24 10.75-24 24v302.74c0 13.25 10.75 24 24 24s24-10.75 24-24V583.6h151.65c68.16 0 123.61-55.45 123.61-123.61s-55.45-123.61-123.61-123.61z m0 199.22H404.37V384.38h151.65c41.69 0 75.61 33.92 75.61 75.61s-33.92 75.61-75.61 75.61z" fill="#FFFFFF" p-id="10174"></path></svg>';
    case 'pdf':
      return '<svg t="1712414525642" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="10337" width="32" height="32"><path d="M658.68 63.95H269.36c-78.06 0-141.33 63.28-141.33 141.33v612.95c0 78.06 63.28 141.33 141.33 141.33h485.49c78.06 0 141.33-63.28 141.33-141.33V301.44c0-31.83-12.64-62.35-35.15-84.85l-117.5-117.5a120.008 120.008 0 0 0-84.85-35.14z" fill="#FF7878" p-id="10338"></path><path d="M861.03 216.59l-117.5-117.5a120.001 120.001 0 0 0-61.75-32.9v131.88c0 41.87 33.94 75.81 75.81 75.81h135.37a120.058 120.058 0 0 0-31.93-57.29zM649.24 554.59c-22.79 0-45.38 3.31-68.16 7.6-26.87-24.76-49.66-54-66.99-86.56 18.5-60.63 19.47-101.76 5.45-121.26-6.43-8.58-17.14-14.04-28.04-14.04-14.02-0.97-27.07 5.46-33.5 17.35-19.47 32.56 8.76 96.5 21.81 122.43-15 46.59-33.5 91.04-57.25 134.32-102.82 44.45-104.97 71.54-104.97 81.29 0 11.89 6.43 23.78 18.31 29.24 4.28 3.12 10.71 4.29 16.16 4.29 27.07 0 58.42-30.41 91.92-90.06 42.26-17.35 84.32-31.39 128.73-41.13 22.79 19.49 50.83 30.41 80.04 32.56 18.5 0 54.14 0 53.75-37.04 1.17-14.04-6.43-37.82-57.26-38.99zM353.53 696.45l-3.26 1.17c9.6-14.06 22.26-25 38.38-31.64-8.44 14.26-20.15 25-35.12 30.47z m132.59-322.36c1.02-1.16 3.17-2.35 4.4-2.35l3.57 1.19c5.52 18.41 5.52 37.8-1.23 56.02-7.96-17.26-11.23-36.64-6.74-54.86zM542.26 584c-23.61 5.5-49.55 13.17-73.16 21.81l-2.45 0.96 0.32-1.94c11.81-23.97 22.64-49.13 32.32-74.28l1.16-2.45 1.17 1.47c12 18.47 27.2 37.51 43.46 53.82l-2.82 0.61z m111.17 13.58c-10.87-0.95-21.54-3.04-32.4-8.35 9.7-2.09 18.24-2.09 27.94-2.09 21.54 0 25.81 5.32 26 8.35-6.41 2.09-13.98 3.23-21.54 2.09z" fill="#F25555" p-id="10339"></path><path d="M648.96 575.14c-9.7 0-18.24 0-27.94 2.09 10.87 5.32 21.54 7.4 32.4 8.35 7.57 1.14 15.14 0 21.54-2.09-0.19-3.03-4.46-8.35-26-8.35z m-147.34-57.57l-1.17-1.47-1.16 2.45c-9.68 25.15-20.52 50.31-32.32 74.28l-0.32 1.94 2.45-0.96c23.61-8.65 49.55-16.31 73.16-21.81l2.81-0.61c-16.25-16.32-31.45-35.35-43.45-53.82z m-7.52-156.65l-3.57-1.19c-1.23 0-3.38 1.19-4.4 2.35-4.5 18.22-1.23 37.6 6.75 54.85 6.74-18.21 6.74-37.59 1.22-56.01z m-143.83 324.7l3.26-1.17c14.97-5.47 26.68-16.21 35.12-30.47-16.12 6.64-28.78 17.58-38.38 31.64z m302.48-67c-29.21-2.14-57.25-13.06-80.04-32.56-44.4 9.75-86.47 23.78-128.73 41.13-33.5 59.65-64.85 90.06-91.92 90.06-5.45 0-11.88-1.17-16.16-4.29-11.88-5.46-18.31-17.35-18.31-29.24 0-9.75 2.14-36.84 104.97-81.29 23.76-43.28 42.26-87.73 57.25-134.32-13.05-25.93-41.29-89.87-21.81-122.43 6.43-11.89 19.47-18.32 33.5-17.35 10.91 0 21.62 5.46 28.04 14.04 14.02 19.49 13.05 60.63-5.45 121.26 17.33 32.56 40.12 61.8 66.99 86.56 22.79-4.29 45.38-7.6 68.16-7.6 50.83 1.17 58.42 24.95 57.25 38.99 0.4 37.04-35.24 37.04-53.74 37.04z" fill="#FFFFFF" p-id="10340"></path></svg>';
    case 'mp3':
    case 'wav':
    case 'aac':
    case 'flac':
    case 'ogg':
    case 'm4a':
    case 'wma':
      return '<svg t="1712415210199" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="12630" width="32" height="32"><path d="M788.14 928.23h-550c-77.32 0-140-62.68-140-140v-550c0-77.32 62.68-140 140-140h550c77.32 0 140 62.68 140 140v550c0 77.32-62.68 140-140 140z" fill="#FF7878" p-id="12631"></path><path d="M682.2 326.8a27.8 27.8 0 0 0-10.04-21.39 27.807 27.807 0 0 0-22.87-5.93l-261.58 49.04c-13.15 2.46-22.67 13.94-22.67 27.32l0.18 250.12a64.668 64.668 0 0 0-22.39-3.98c-35.76 0-64.74 28.99-64.74 64.74 0 35.76 28.99 64.74 64.74 64.74s64.74-28.99 64.74-64.74V450.77l232.29-43.55v169.7a64.668 64.668 0 0 0-22.39-3.98c-35.76 0-64.74 28.99-64.74 64.74s28.99 64.74 64.74 64.74 64.74-28.99 64.74-64.74c-0.01-1.19-0.01-310.88-0.01-310.88z" fill="#EF5252" p-id="12632"></path><path d="M682.2 314.8a27.8 27.8 0 0 0-10.04-21.39 27.807 27.807 0 0 0-22.87-5.93l-261.58 49.04c-13.15 2.46-22.67 13.94-22.67 27.32l0.18 250.12a64.668 64.668 0 0 0-22.39-3.98c-35.76 0-64.74 28.99-64.74 64.74 0 35.76 28.99 64.74 64.74 64.74s64.74-28.99 64.74-64.74V438.77l232.29-43.55v169.7a64.668 64.668 0 0 0-22.39-3.98c-35.76 0-64.74 28.99-64.74 64.74s28.99 64.74 64.74 64.74 64.74-28.99 64.74-64.74c-0.01-1.19-0.01-310.88-0.01-310.88z" fill="#FFFFFF" p-id="12633"></path></svg>';
    case 'mp4':
    case 'avi':
    case 'mov':
    case 'wmv':
    case 'flv':
    case 'mkv':
    case 'm4v':
    case 'mpg':
    case 'mpeg':
    case '3gp':
    case '3g2':
    case 'webm':
      return '<svg t="1712415228957" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="12798" width="32" height="32"><path d="M65.3 755.4V269.9c0-78.1 63.3-141.3 141.3-141.3h612.9c78.1 0 141.3 63.3 141.3 141.3v485.5c0 78.1-63.3 141.3-141.3 141.3H206.7c-78.1 0-141.4-63.3-141.4-141.3z" fill="#8B72F7" p-id="12799"></path><path d="M557.3 406.6H308c-22.1 0-40 17.9-40 40v155.9c0 22.1 17.9 40 40 40h249.3c22.1 0 40-17.9 40-40V446.6c0-22-17.9-40-40-40zM738.8 412.7L655 491c-19.5 18.2-19.5 49 0 67.2l83.8 78.3c14.2 13.3 37.5 3.2 37.5-16.2V428.9c0-19.4-23.3-29.5-37.5-16.2z" fill="#7463EA" p-id="12800"></path><path d="M557.3 394.6H308c-22.1 0-40 17.9-40 40v155.9c0 22.1 17.9 40 40 40h249.3c22.1 0 40-17.9 40-40V434.6c0-22-17.9-40-40-40zM738.8 400.7L655 479c-19.5 18.2-19.5 49 0 67.2l83.8 78.3c14.2 13.3 37.5 3.2 37.5-16.2V416.9c0-19.4-23.3-29.5-37.5-16.2z" fill="#FFFFFF" p-id="12801"></path></svg>';
    case 'zip':
    case 'rar':
    case '7z':
    case 'tar':
    case 'gz':
    case 'bz2':
      return '<svg t="1712414898679" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="12462" width="32" height="32"><path d="M659.72 64.81H270.4c-78.06 0-141.33 63.28-141.33 141.33v612.95c0 78.06 63.28 141.33 141.33 141.33h485.49c78.06 0 141.33-63.28 141.33-141.33V302.3c0-31.83-12.64-62.35-35.15-84.85l-117.5-117.5a120.008 120.008 0 0 0-84.85-35.14z" fill="#FFC757" p-id="12463"></path><path d="M304.85 196.81h71.52c7.86 0 14.24-6.38 14.24-14.24v-71.52c0-7.87-6.38-14.24-14.24-14.24h-71.52c-7.86 0-14.24 6.38-14.24 14.24v71.52c0 7.86 6.38 14.24 14.24 14.24zM476.37 196.81h-71.52c-7.86 0-14.24 6.38-14.24 14.24v71.52c0 7.87 6.38 14.24 14.24 14.24h71.52c7.86 0 14.24-6.38 14.24-14.24v-71.52c0-7.86-6.38-14.24-14.24-14.24zM304.85 396.81h71.52c7.86 0 14.24-6.38 14.24-14.24v-71.52c0-7.87-6.38-14.24-14.24-14.24h-71.52c-7.86 0-14.24 6.38-14.24 14.24v71.52c0 7.86 6.38 14.24 14.24 14.24zM476.37 396.81h-71.52c-7.86 0-14.24 6.38-14.24 14.24v71.52c0 7.87 6.38 14.24 14.24 14.24h71.52c7.86 0 14.24-6.38 14.24-14.24v-71.52c0-7.86-6.38-14.24-14.24-14.24zM455.65 552.02H321.52c-17.07 0-30.91 13.84-30.91 30.91v155.48c0 17.07 13.84 30.91 30.91 30.91h134.13c17.07 0 30.91-13.84 30.91-30.91V582.93c0-17.07-13.84-30.91-30.91-30.91z m-4.93 164.59c0 13.72-11.12 24.84-24.84 24.84h-74.6c-13.72 0-24.84-11.12-24.84-24.84v-27.36c0-13.72 11.12-24.84 24.84-24.84h74.6c13.72 0 24.84 11.12 24.84 24.84v27.36zM862.07 217.45l-117.5-117.5a120.001 120.001 0 0 0-61.75-32.9v131.88c0 41.87 33.94 75.81 75.81 75.81H894a120.021 120.021 0 0 0-31.93-57.29z" fill="#F79F2B" p-id="12464"></path><path d="M304.85 184.81h71.52c7.86 0 14.24-6.38 14.24-14.24V99.05c0-7.87-6.38-14.24-14.24-14.24h-71.52c-7.86 0-14.24 6.38-14.24 14.24v71.52c0 7.86 6.38 14.24 14.24 14.24zM476.37 184.81h-71.52c-7.86 0-14.24 6.38-14.24 14.24v71.52c0 7.87 6.38 14.24 14.24 14.24h71.52c7.86 0 14.24-6.38 14.24-14.24v-71.52c0-7.86-6.38-14.24-14.24-14.24zM304.85 384.81h71.52c7.86 0 14.24-6.38 14.24-14.24v-71.52c0-7.87-6.38-14.24-14.24-14.24h-71.52c-7.86 0-14.24 6.38-14.24 14.24v71.52c0 7.86 6.38 14.24 14.24 14.24zM476.37 384.81h-71.52c-7.86 0-14.24 6.38-14.24 14.24v71.52c0 7.87 6.38 14.24 14.24 14.24h71.52c7.86 0 14.24-6.38 14.24-14.24v-71.52c0-7.86-6.38-14.24-14.24-14.24zM455.65 540.02H321.52c-17.07 0-30.91 13.84-30.91 30.91v155.48c0 17.07 13.84 30.91 30.91 30.91h134.13c17.07 0 30.91-13.84 30.91-30.91V570.93c0-17.07-13.84-30.91-30.91-30.91z m-4.93 164.59c0 13.72-11.12 24.84-24.84 24.84h-74.6c-13.72 0-24.84-11.12-24.84-24.84v-27.36c0-13.72 11.12-24.84 24.84-24.84h74.6c13.72 0 24.84 11.12 24.84 24.84v27.36z" fill="#FFFFFF" p-id="12465"></path></svg>';
    default:
      return '<svg t="1712414453009" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="9667" width="32" height="32"><path d="M659.72 64.19H270.4c-78.06 0-141.33 63.28-141.33 141.33v612.95c0 78.06 63.28 141.33 141.33 141.33h485.49c78.06 0 141.33-63.28 141.33-141.33V301.69c0-31.83-12.64-62.35-35.15-84.85l-117.5-117.5a120.017 120.017 0 0 0-84.85-35.15z" fill="#C7DADD" p-id="9668"></path><path d="M862.07 216.84l-117.5-117.5a120.001 120.001 0 0 0-61.75-32.9v131.88c0 41.87 33.94 75.81 75.81 75.81H894a119.975 119.975 0 0 0-31.93-57.29zM512.53 295.92c-82.1 0-148.9 66.8-148.9 148.9 0 12.53 10.15 22.68 22.68 22.68s22.68-10.15 22.68-22.68c0-57.09 46.45-103.54 103.54-103.54 27.63 0 53.62 10.78 73.19 30.35s30.35 45.57 30.35 73.19c0 21.11-6.31 41.41-18.24 58.71-11.68 16.92-27.89 29.89-46.9 37.48-37.11 14.83-61.08 49.9-61.08 89.33V632c0 12.53 10.15 22.68 22.68 22.68s22.68-10.15 22.68-22.68v-1.66c0-20.78 12.78-39.31 32.56-47.21 27.33-10.92 50.63-29.54 67.4-53.84 17.18-24.9 26.26-54.11 26.26-84.47 0-39.74-15.5-77.13-43.64-105.27-28.13-28.13-65.52-43.63-105.26-43.63z" fill="#9DC0C9" p-id="9669"></path><path d="M500.973351 750.25961a30.24 30.24 0 1 0 23.144694-55.876234 30.24 30.24 0 1 0-23.144694 55.876234Z" fill="#9DC0C9" p-id="9670"></path><path d="M618.4 327.31c-28.14-28.14-65.52-43.64-105.26-43.64-82.1 0-148.9 66.8-148.9 148.9 0 12.53 10.15 22.68 22.68 22.68s22.68-10.15 22.68-22.68c0-57.09 46.45-103.54 103.54-103.54 27.63 0 53.62 10.78 73.19 30.35s30.35 45.57 30.35 73.19c0 21.11-6.31 41.41-18.24 58.71-11.68 16.92-27.89 29.89-46.9 37.48-37.11 14.83-61.08 49.9-61.08 89.33v1.66c0 12.53 10.15 22.68 22.68 22.68s22.68-10.15 22.68-22.68v-1.66c0-20.78 12.78-39.31 32.56-47.21 27.33-10.92 50.63-29.54 67.4-53.84 17.18-24.9 26.26-54.11 26.26-84.47 0-39.74-15.5-77.12-43.64-105.26z" fill="#FFFFFF" p-id="9671"></path><path d="M513.14 710.09m-30.24 0a30.24 30.24 0 1 0 60.48 0 30.24 30.24 0 1 0-60.48 0Z" fill="#FFFFFF" p-id="9672"></path></svg>';
  }
};

const downloadFiles = async (Id, CompanyId, fileName) => {
  try {
    const response = await axios.post('http://47.97.176.174:8087/chainmaker?cmb=GetFileById', {
      Id: Id,
      CompanyId: CompanyId,
      FileName: fileName
    }, {
      responseType: 'blob'
    });
    const blobType = response.headers['content-type'] || 'application/octet-stream';
    const blob = new Blob([response.data], { type: blobType });
    const downloadUrl = URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = downloadUrl;
    link.setAttribute('download', fileName);
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
    URL.revokeObjectURL(downloadUrl);
    ElMessage.success('下载成功');
  } catch (error) {
    console.error('下载文件失败:', error);
    ElMessage.error('下载失败');
  }
}

const downloadFile = async (index, row) => {
  await axios.post('http://47.97.176.174:8087/chainmaker?cmb=GetPNGImageById', {
    Id: row.Id,
    CompanyId: row.CompanyId
  }, {
    responseType: 'blob'
  })
      .then(function (response) {
        const blob = new Blob([response.data], {type: 'application/pdf'})
        const downloadUrl = URL.createObjectURL(blob)
        const link = document.createElement('a')
        link.href = downloadUrl
        link.setAttribute('download', '碳排放报告.pdf')
        document.body.appendChild(link)
        link.click()
        document.body.removeChild(link)
        URL.revokeObjectURL(downloadUrl)
        ElMessage.success('下载成功');
      })
      .catch(function (error) {
        console.log(error)
        ElMessage.error('下载失败');
      })
}

const handleCheck = async (index, row) => {
  if (row.Type === "0") {
    await axios.post('http://47.97.176.174:8087/chainmaker?cmb=DAPExamineForPGE', {
      Id : row.Id,
      State : 1,
      DAPadmin : 'DAPAdmin1'
    })
        .then(function (response) {
          refresh()
          ElMessage.success('审核成功')
          console.log(response)
        })
        .catch(function (error) {
          ElMessage.error('审核失败')
          console.log(error)
        })
  } else if (row.Type === "1") {
    await axios.post('http://47.97.176.174:8087/chainmaker?cmb=DAPExamineForCE', {
      Id : row.Id,
      State : 1,
      DAPadmin : 'DAPAdmin1'
    })
        .then(function (response) {
          refresh()
          ElMessage.success('审核成功')
          console.log(response)
        })
        .catch(function (error) {
          ElMessage.error('审核失败')
          console.log(error)
        })
  } else {

  }
  console.log(index, row)
}
const handleClose = async (index, row) => {
  if (row.Type === "0") {
    await axios.post('http://47.97.176.174:8087/chainmaker?cmb=DAPExamineForPGE', {
      Id : row.Id,
      State : 2,
      DAPadmin : 'DAPAdmin1'
    })
        .then(function (response) {
          refresh()
          ElMessage.success('审核成功')
          console.log(response)
        })
        .catch(function (error) {
          ElMessage.error('审核失败')
          console.log(error)
        })
  } else if (row.Type === "1") {
    await axios.post('http://47.97.176.174:8087/chainmaker?cmb=DAPExamineForCE', {
      Id : row.Id,
      State : 2,
      DAPadmin : 'DAPAdmin1'
    })
        .then(function (response) {
          refresh()
          ElMessage.success('审核成功')
          console.log(response)
        })
        .catch(function (error) {
          ElMessage.error('审核失败')
          console.log(error)
        })
  } else {

  }
  console.log(index, row)
}
const handleInfo = async (index, row) => {
  if (row.Type === "0") {
    currentCompanyData.value = await axios.post('http://47.97.176.174:8087/chainmaker?cmb=SearchPGEById', {
      Id : row.Id
    })
        .then(function (response) {
          console.log(response)
          return response.data.data
        })
        .catch(function (error) {
          console.log(error);
        })
    dialogVisibleForPGE.value = true
  } else if (row.Type === "1") {
    currentCompanyData.value = await axios.post('http://47.97.176.174:8087/chainmaker?cmb=SearchCEById', {
      Id : row.Id
    })
        .then(function (response) {
          console.log(response)
          return response.data.data
        })
        .catch(function (error) {
          console.log(error);
        })
    dialogVisibleForCE.value = true
  } else {

  }
  console.log(index, row)
}

const handleSizeChange = (val) => {
  console.log(`${val} items per page`)
}
const handleCurrentChange = (val) => {
  console.log(`current page: ${val}`)
}


const translateCalculatedDataSource = (code) => {
  const dataSourceMap = {
    0: '检测值',
    1: '计算值'
  };
  return dataSourceMap[code] || '';
};

const translateDefaultDataSource = (code) => {
  const dataSourceMap = {
    0: '检测值',
    1: '缺省值'
  };
  return dataSourceMap[code] || '';
};

const translateExtendedDataSource = (code) => {
  const dataSourceMap = {
    0: '检测值',
    1: '化学计算',
    2: '缺省值'
  };
  return dataSourceMap[code] || '';
};
</script>

<template>
  <el-scrollbar>
    <div class="main">
      <el-card shadow="never">
        <div style="margin: 10px 10px 0 10px">
          <el-row :gutter="12">
            <el-col :span="6">
              <div class="statistic-card" style="background-color: #dedfe0">
                <el-statistic :value="totalCount" >
                  <template #title>
                    <div style="display: inline-flex; align-items: center">
                      Total
                    </div>
                  </template>
                </el-statistic>
                <div class="statistic-footer">
                  <div class="footer-item">
                    <span>than yesterday</span>
                    <span class="green">
              10%
              <el-icon>
                <CaretTop />
              </el-icon>
            </span>
                  </div>
                </div>
              </div>
            </el-col>
            <el-col :span="6">
              <div class="statistic-card" style="background-color: #fcd3d3">
                <el-statistic :value="unreviewedCount">
                  <template #title>
                    <div style="display: inline-flex; align-items: center">
                      未审核
                      <el-tooltip
                          effect="dark"
                          content="未审核"
                          placement="top"
                      >
                        <el-icon style="margin-left: 4px" :size="12">
                          <CircleClose />
                        </el-icon>
                      </el-tooltip>
                    </div>
                  </template>
                </el-statistic>
                <div class="statistic-footer">
                  <div class="footer-item">
                    <span>than yesterday</span>
                    <span class="red">
              12%
              <el-icon>
                <CaretBottom />
              </el-icon>
            </span>
                  </div>
                </div>
              </div>
            </el-col>
            <el-col :span="6">
              <div class="statistic-card" style="background-color: #d1edc4;">
                <el-statistic :value="reviewedPassedCount">
                  <template #title>
                    <div style="display: inline-flex; align-items: center">
                      已审核通过
                      <el-tooltip
                          effect="dark"
                          content="已审核通过"
                          placement="top"
                      >
                        <el-icon style="margin-left: 4px" :size="12">
                          <CircleCheck />
                        </el-icon>
                      </el-tooltip>
                    </div>
                  </template>
                </el-statistic>
                <div class="statistic-footer">
                  <div class="footer-item">
                    <span>than yesterday</span>
                    <span class="green">
              24%
              <el-icon>
                <CaretTop />
              </el-icon>
            </span>
                  </div>
                </div>
              </div>
            </el-col>
            <el-col :span="6">
              <div class="statistic-card" style="background-color: #f8e3c5">
                <el-statistic :value="reviewedNotPassedCount">
                  <template #title>
                    <div style="display: inline-flex; align-items: center">
                      未审核通过
                      <el-tooltip
                          effect="dark"
                          content="未审核通过"
                          placement="top"
                      >
                        <el-icon style="margin-left: 4px" :size="12">
                          <Warning />
                        </el-icon>
                      </el-tooltip>
                    </div>
                  </template>
                </el-statistic>
                <div class="statistic-footer">
                  <div class="footer-item">
                    <span>than yesterday</span>
                    <span class="green">
              16%
              <el-icon>
                <CaretTop />
              </el-icon>
            </span>
                  </div>
                </div>
              </div>
            </el-col>
          </el-row>
        </div>
        <el-divider style="margin: 10px auto" />
        <div class="main-top">
          <div>
            <el-tooltip content="刷新" placement="top">
              <el-button color="#40485b" :icon="Refresh" @click="refresh">刷新</el-button>
            </el-tooltip>
            <el-button type="success" :icon="Plus" disabled>添加</el-button>
            <el-button type="primary" :icon="Edit" disabled>修改</el-button>
            <el-button type="danger" :icon="Delete" disabled>删除</el-button>
            <el-tooltip content="导出为Excel" placement="top">
              <el-button type="primary" :icon="Download" @click="exportToExcel">导出</el-button>
            </el-tooltip>
          </div>
          <div>
            <el-button :icon="Search" circle></el-button>
            <el-input v-model="searchQuery" placeholder="请输入公司ID" style="width: 200px; margin-left: 10px;" clearable></el-input>
          </div>
        </div>
        <div class="main-middle">
          <el-table :data="tableData.slice((currentPage - 1) * pageSize, currentPage * pageSize)" border style="width: 100%" table-layout="auto" :header-cell-style="{ background:'#f4f4f5' }">
            <el-table-column prop="Id" label="ID" align="center" />
            <el-table-column prop="CompanyId" label="公司ID" align="center" />
            <el-table-column prop="createAt" label="创建时间" align="center" sortable />
            <el-table-column prop="updateAt" label="更新时间" align="center" sortable />
            <el-table-column prop="State" label="审核状态" align="center"
                             :filters="stateOptions"
                             :filter-method="filterState"
                             filter-placement="bottom"
            >
              <template #default="scope">
                <el-tag :type="stateTagType(scope.row.State)" :disable-transitions="false">
                  {{ stateTagText(scope.row.State) }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column prop="Type" label="企业类型" align="center"
                             :filters="typeOptions"
                             :filter-method="filterType"
                             filter-placement="bottom"
            >
              <template #default="scope">
                <el-tag :type="typeTagType(scope.row.Type)" :disable-transitions="false">
                  {{ typeTagText(scope.row.Type) }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column label="证明材料" align="center">
              <template #default="scope">
                <el-button size="small" type="primary" @click="showFiles(scope.$index, scope.row)">查看</el-button>
              </template>
            </el-table-column>
            <el-table-column label="碳排放报告" align="center">
              <template #default="scope">
                <el-button size="small" type="primary" @click="downloadFile(scope.$index, scope.row)">下载</el-button>
              </template>
            </el-table-column>
            <el-table-column label="操作" align="center">
              <template #default="scope">
                <el-button size="small" type="success" @click="handleCheck(scope.$index, scope.row)" :icon="CircleCheckFilled"
                >通过</el-button
                >
                <el-button
                    size="small"
                    type="danger"
                    @click="handleClose(scope.$index, scope.row)"
                    :icon="CircleCloseFilled"
                >不通过</el-button
                >
                <el-button size="small" type="primary" @click="handleInfo(scope.$index, scope.row)" :icon="InfoFilled">详情</el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>
        <el-divider style="margin: 10px auto" />
        <el-pagination
            v-model:current-page="currentPage"
            v-model:page-size="pageSize"
            :page-sizes="[15, 20, 25, 30]"
            :small="small"
            :disabled="disabled"
            :background="background"
            layout="total, sizes, ->, prev, pager, next, jumper"
            :total="tableData.length"
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
        />
      </el-card>
    </div>
    <el-dialog
        v-model="dialogVisibleForPGE"
        title="电网企业"
        width="60%"
    >
      <span>{{ currentCompanyData.Data.CompanyId }}</span>
      <template #footer>
        <div class="dialog-footer">
          <el-tabs type="border-card">
            <el-tab-pane label="附表1 报告主体2024年二氧化碳排放量报告">
              <el-descriptions
                  :column="1"
                  border
              >
                <el-descriptions-item label="企业二氧化碳排放总量(tCO₂)" label-align="center" align="center">
                  {{ currentCompanyData.Total }}
                </el-descriptions-item>
                <el-descriptions-item align="center">
                  <template #label>
                    <div class="cell-item">
                      使用六氟化硫设备修理与退役过程产生的排放(tCO₂)
                    </div>
                  </template>
                  {{ currentCompanyData.FSCO }}
                </el-descriptions-item>
                <el-descriptions-item align="center">
                  <template #label>
                    <div class="cell-item">
                      输配电引起的二氧化碳排放(tCO₂)
                    </div>
                  </template>
                  {{ currentCompanyData.SPDCO }}
                </el-descriptions-item>
              </el-descriptions>
            </el-tab-pane>
            <el-tab-pane label="附表2 报告主体活动水平数据">
              <el-descriptions
                  title="六氟化硫回收*"
                  direction="vertical"
                  :column="1"
              >
                <el-descriptions-item>
                  <el-table :data="currentCompanyData.Data.PGERepairREC" border :header-cell-style="{ background:'#f4f4f5' }">
                    <el-table-column prop="Id" label="修理设备" align="center" />
                    <el-table-column prop="RECV" label="设备容量(千克)" align="center" />
                    <el-table-column prop="RECR" label="实际回收量(千克)" align="center" />
                  </el-table>
                </el-descriptions-item>
                <el-descriptions-item>
                  <el-table :data="currentCompanyData.Data.PGERetireREC" border :header-cell-style="{ background:'#f4f4f5' }">
                    <el-table-column prop="Id" label="退役设备" align="center" />
                    <el-table-column prop="RECV" label="设备容量(千克)" align="center" />
                    <el-table-column prop="RECR" label="实际回收量(千克)" align="center" />
                  </el-table>
                </el-descriptions-item>
              </el-descriptions>
              <el-descriptions
                  title="输配电损失"
                  :column="1"
                  border
              >
                <el-descriptions-item align="center">
                  <template #label>
                    <div class="cell-item">
                      电厂上网电量(兆瓦时)
                    </div>
                  </template>
                  {{ currentCompanyData.Data.ELSW }}
                </el-descriptions-item>
                <el-descriptions-item align="center">
                  <template #label>
                    <div class="cell-item">
                      自外省输入电量(兆瓦时)
                    </div>
                  </template>
                  {{ currentCompanyData.Data.ELSR }}
                </el-descriptions-item>
                <el-descriptions-item align="center">
                  <template #label>
                    <div class="cell-item">
                      向外省输出电量(兆瓦时)
                    </div>
                  </template>
                  {{ currentCompanyData.Data.ELSC }}
                </el-descriptions-item>
                <el-descriptions-item align="center">
                  <template #label>
                    <div class="cell-item">
                      售电量(兆瓦时)
                    </div>
                  </template>
                  {{ currentCompanyData.Data.ELSD }}
                </el-descriptions-item>
                <el-descriptions-item align="center">
                  <template #label>
                    <div class="cell-item">
                      输配电量(兆瓦时)
                    </div>
                  </template>
                  {{ currentCompanyData.Data.ELSPDL }}
                </el-descriptions-item>
              </el-descriptions>
            </el-tab-pane>
            <el-tab-pane label="附表3 报告主体排放因子">
              <el-descriptions
                  title="输配电损失"
                  direction="vertical"
                  :column="3"
                  border
              >
                <el-descriptions-item align="center">
                  <template #label>
                    <div class="cell-item">
                    </div>
                  </template>
                  电力
                </el-descriptions-item>
                <el-descriptions-item label="数据" label-align="center" align="center">
                  {{ currentCompanyData.SPDCO }}
                </el-descriptions-item>
                <el-descriptions-item label="单位" label-align="center" align="center">
                  吨二氧化碳/兆瓦时
                </el-descriptions-item>
              </el-descriptions>
            </el-tab-pane>
          </el-tabs>
        </div>
      </template>
    </el-dialog>
    <el-dialog
        v-model="dialogVisibleForCE"
        title="化工企业"
        width="70%"
    >
      <span>{{ currentCompanyData.Data.CompanyId }}</span>
      <template #footer>
        <div class="dialog-footer">
          <el-tabs type="border-card">
            <el-tab-pane label="附表1">
              <el-descriptions
                  title="报告主体2024年温室气体排放量汇总"
                  direction="vertical"
                  :column="1"
              >
                <el-descriptions-item>
                  <el-table :data="currentCompanyData.Data.COSummaryS" border :header-cell-style="{ background:'#f4f4f5' }">
                    <el-table-column prop="Breed" label="源类别" align="center" />
                    <el-table-column prop="WS" label="温室气体本身质量（单位：吨）" align="center" />
                    <el-table-column prop="CO" label="CO₂当量（单位：吨CO₂当量）" align="center" />
                  </el-table>
                </el-descriptions-item>
              </el-descriptions>
            </el-tab-pane>
            <el-tab-pane label="附表2">
              <el-descriptions
                  title="化石燃料燃烧的活动水平和排放因子数据一览表"
                  direction="vertical"
                  :column="1"
              >
                <el-descriptions-item>
                  <el-table :data="currentCompanyData.Data.Ecors" border :header-cell-style="{ background:'#f4f4f5' }">
                    <el-table-column prop="Breed" label="燃料品种" align="center" />
                    <el-table-column prop="AD" label="燃烧量（吨或万Nm³）" align="center" />
                    <el-table-column prop="CC" label="含碳量（tC/吨或tC/万Nm³）" align="center" />
                    <el-table-column prop="CCDS" label="数据来源" align="center">
                      <template #default="{ row }">
                        {{ translateCalculatedDataSource(row.CCDS) }}
                      </template>
                    </el-table-column>
                    <el-table-column prop="NCV" label="低位发热量*（GJ/吨或GJ/万Nm³）" align="center" />
                    <el-table-column prop="NCVDS" label="数据来源" align="center">
                      <template #default="{ row }">
                        {{ translateDefaultDataSource(row.NCVDS) }}
                      </template>
                    </el-table-column>
                    <el-table-column prop="EF" label="单位热值含碳量*（tC/GJ）" align="center" />
                    <el-table-column prop="OF" label="碳氧化率（%）" align="center" />
                    <el-table-column prop="OFDS" label="数据来源" align="center">
                      <template #default="{ row }">
                        {{ translateDefaultDataSource(row.OFDS) }}
                      </template>
                    </el-table-column>
                  </el-table>
                </el-descriptions-item>
              </el-descriptions>
            </el-tab-pane>
            <el-tab-pane label="附表3">
              <el-descriptions
                  title="工业生产过程CO₂排放的活动水平和排放因子数据一览表"
                  direction="vertical"
                  :column="1"
              >
                <el-descriptions-item>
                  <el-table :data="currentCompanyData.Data.Eghg.GYSCCORC.GYSCCOR" border :header-cell-style="{ background:'#f4f4f5' }">
                    <el-table-column label="碳输入" align="center">
                      <el-table-column prop="Breed" label="物料名称" align="center" />
                      <el-table-column prop="AD" label="活动水平数据（单位：吨或万Nm³）" align="center" />
                      <el-table-column prop="CC" label="含碳量（单位：tC/吨）" align="center" />
                      <el-table-column prop="CCDS" label="数据来源" align="center">
                        <template #default="{ row }">
                          {{ translateExtendedDataSource(row.CCDS) }}
                        </template>
                      </el-table-column>
                    </el-table-column>
                  </el-table>
                </el-descriptions-item>
                <el-descriptions-item>
                  <el-table :data="currentCompanyData.Data.Eghg.GYSCCORC.GYSCCOC" border :header-cell-style="{ background:'#f4f4f5' }">
                    <el-table-column label="碳输出" align="center">
                      <el-table-column prop="Breed" label="物料名称" align="center" />
                      <el-table-column prop="AD" label="活动水平数据（单位：吨或万Nm³）" align="center" />
                      <el-table-column prop="CC" label="含碳量（单位：tC/吨）" align="center" />
                      <el-table-column prop="CCDS" label="数据来源" align="center">
                        <template #default="{ row }">
                          {{ translateExtendedDataSource(row.CCDS) }}
                        </template>
                      </el-table-column>
                    </el-table-column>
                  </el-table>
                </el-descriptions-item>
              </el-descriptions>
            </el-tab-pane>
            <el-tab-pane label="附表4">
              <el-descriptions
                  title="碳酸盐使用的活动水平和排放因子数据一览表"
                  direction="vertical"
                  :column="1"
              >
                <el-descriptions-item>
                  <el-table :data="currentCompanyData.Data.Eghg.CACOS" border :header-cell-style="{ background:'#f4f4f5' }">
                    <el-table-column prop="Breed" label="碳酸盐种类" align="center" />
                    <el-table-column prop="AD" label="消耗量（单位：吨）" align="center" />
                    <el-table-column prop="EF" label="CO₂排放因子（吨CO₂/吨碳酸盐）" align="center" />
                    <el-table-column prop="EFDS" label="数据来源" align="center">
                      <template #default="{ row }">
                        {{ translateExtendedDataSource(row.EFDS) }}
                      </template>
                    </el-table-column>
                  </el-table>
                </el-descriptions-item>
              </el-descriptions>
            </el-tab-pane>
            <el-tab-pane label="附表5">
              <el-descriptions
                  title="硝酸生产过程的活动水平和N₂O排放因子数据一览表"
                  direction="vertical"
                  :column="1"
              >
                <el-descriptions-item>
                  <el-table :data="currentCompanyData.Data.Eghg.XSCOS" border :header-cell-style="{ background:'#f4f4f5' }">
                    <el-table-column prop="Breed" label="硝酸生产工艺类型" align="center" />
                    <el-table-column prop="AD" label="硝酸产量（吨）" align="center" />
                    <el-table-column prop="EF" label="N₂O生成因子（kgN₂O/吨硝酸）" align="center" />
                    <el-table-column prop="EFDS" label="数据来源" align="center">
                      <template #default="{ row }">
                        {{ translateDefaultDataSource(row.EFDS) }}
                      </template>
                    </el-table-column>
                    <el-table-column prop="NK" label="N₂O去除率（%）" align="center" />
                    <el-table-column prop="NKDS" label="数据来源" align="center">
                      <template #default="{ row }">
                        {{ translateDefaultDataSource(row.NKDS) }}
                      </template>
                    </el-table-column>
                    <el-table-column prop="UK" label="尾气处理设备使用率（%）" align="center" />
                    <el-table-column prop="UKDS" label="数据来源" align="center">
                      <template #default="{ row }">
                        {{ translateDefaultDataSource(row.UKDS) }}
                      </template>
                    </el-table-column>
                  </el-table>
                </el-descriptions-item>
              </el-descriptions>
            </el-tab-pane>
            <el-tab-pane label="附表6">
              <el-descriptions
                  title="己二酸生产过程的活动水平和N₂O排放因子数据一览表"
                  direction="vertical"
                  :column="1"
              >
                <el-descriptions-item>
                  <el-table :data="currentCompanyData.Data.Eghg.YRSCOS" border :header-cell-style="{ background:'#f4f4f5' }">
                    <el-table-column prop="Breed" label="己二酸生产工艺类型" align="center" />
                    <el-table-column prop="AD" label="己二酸产量（吨）" align="center" />
                    <el-table-column prop="EF" label="N₂O生成因子（kgN₂O/吨己二酸）" align="center" />
                    <el-table-column prop="EFDS" label="数据来源" align="center">
                      <template #default="{ row }">
                        {{ translateDefaultDataSource(row.EFDS) }}
                      </template>
                    </el-table-column>
                    <el-table-column prop="NK" label="N₂O去除率（%）" align="center" />
                    <el-table-column prop="NKDS" label="数据来源" align="center">
                      <template #default="{ row }">
                        {{ translateDefaultDataSource(row.NKDS) }}
                      </template>
                    </el-table-column>
                    <el-table-column prop="UK" label="尾气处理设备使用率（%）" align="center" />
                    <el-table-column prop="UKDS" label="数据来源" align="center">
                      <template #default="{ row }">
                        {{ translateDefaultDataSource(row.UKDS) }}
                      </template>
                    </el-table-column>
                  </el-table>
                </el-descriptions-item>
              </el-descriptions>
            </el-tab-pane>
            <el-tab-pane label="附表7">
              <el-descriptions
                  title="净购入的电力和热力消费活动水平和排放因子数据一览表"
                  direction="vertical"
                  :column="1"
              >
                <el-descriptions-item>
                  <el-table :data="currentCompanyData.Data.Drcos" border :header-cell-style="{ background:'#f4f4f5' }">
                    <el-table-column prop="Breed" label="类型" align="center" />
                    <el-table-column prop="AD" label="净购入量（MWh或GJ）" align="center" />
                    <el-table-column prop="AD1" label="购入量（MWh或GJ）" align="center" />
                    <el-table-column prop="AD2" label="外供量（MWh或GJ）" align="center" />
                    <el-table-column prop="EF" label="CO₂排放因子（tCO₂/MWh或tCO₂/GJ）" align="center" />
                  </el-table>
                </el-descriptions-item>
              </el-descriptions>
            </el-tab-pane>
          </el-tabs>
        </div>
      </template>
    </el-dialog>
    <el-dialog
        v-model="dialogVisibleForFiles"
        title="证明材料"
        width="30%"
    >
      <div class="file-list-container">
        <div v-for="(fileName, index) in fileList" :key="index" class="file-item">
          <div class="file-info">
            <span class="file-icon" v-html="fileIcon(fileName)"></span>
            <span class="file-name">{{ fileName }}</span>
          </div>
          <el-button size="small" color="#FFFFFF" :icon="Download" circle @click="downloadFiles(currentRow.Id, currentRow.CompanyId, fileName)"></el-button>
        </div>
      </div>
    </el-dialog>
  </el-scrollbar>
</template>

<style scoped>
.main {
  margin: 15px;
}

.main-top {
  margin: 0 5px 10px 5px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.main-middle {
  margin: 5px;
}

.cell-item {
  display: flex;
  justify-content: center;
  align-items: center;
}

:deep(.el-card__body) {
  padding: 0;
}

:global(h2#card-usage ~ .example .example-showcase) {
  background-color: var(--el-fill-color) !important;
}

.el-statistic {
  --el-statistic-content-font-size: 28px;
}

.statistic-card {
  padding: 20px;
  border-radius: 4px;
  background-color: var(--el-bg-color-overlay);
  transition: all .3s ease;
  background-image: url("@/assets/dataExamine/starTrails.svg");
  background-size: cover;
  background-repeat: no-repeat;
  background-position-x: 150px;
  background-position-y: -225px;
}

.statistic-card:hover {
  transform: translateY(-4px) scale(1.02);
  box-shadow: 0 14px 24px #0003;
}

.statistic-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  font-size: 12px;
  color: var(--el-text-color-regular);
}

.statistic-footer .footer-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.statistic-footer .footer-item span:last-child {
  display: inline-flex;
  align-items: center;
  margin-left: 4px;
}

.green {
  color: var(--el-color-success);
}
.red {
  color: var(--el-color-error);
}

.el-pagination {
  margin: 10px;
}

.file-list-container {
  max-height: 300px;
  overflow-y: auto;
}

.file-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 10px 0;
  padding: 10px;
  border-radius: 10px;
  background-color: #f4f4f5;
}

.file-icon {
  display: flex;
  justify-content: center;
  align-items: center;
}

.file-info {
  display: flex;
  justify-content: center;
  align-items: center;
}

.file-name {
  margin: 0 10px;
  font-weight: 500;
}
</style>
