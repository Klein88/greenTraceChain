// import './assets/main.css'
import './assets/base.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import Vuex from 'vuex';
import ElementPlus, {ElMessage} from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
const app = createApp(App)

const store = new Vuex.Store({
    state: {
        companyData:null,
        FormData:null,
        formData2:null,
        PGERepairREC:null,
        PGERetireREC:null,
        greenhouse:null,
        FossilFuel:null,
        IndustryInput:null,
        IndustryOutput:null,
        M2CO3:null,
        HNO3:null,
        HOO:null,
        Electrical:null,
    },
    mutations: {
        setCompanyData(state, companyData) {
            state.companyData = companyData;
        },
        // 添加更新 FormData 的 mutation
        updateFormData(state, newFormData) {
            state.FormData = newFormData;
        },
        updateFormData2(state, newFormData) {
            state.formData2 = newFormData;
        },
        // 添加更新 tableData 的 mutation
        updatePGERepairREC(state, newTableData) {
            state.PGERepairREC = newTableData;
        },
        updatePGERetireREC(state, newTableData) {
            state.PGERetireREC = newTableData;
        },
        updategreenhouse(state,newgreenhouse){
            state.greenhouse = newgreenhouse;
        },
        updateFossilFuel(state,newFossilFuel){
            state.FossilFuel = newFossilFuel;
        },
        updateIndustryInput(state,newIndustryInput){
            state.IndustryInput = newIndustryInput;
        },
        updateIndustryOutput(state,newIndustryOutput){
            state.IndustryOutput = newIndustryOutput;
        },
        updateM2CO3(state,newM2CO3){
            state.M2CO3 = newM2CO3;
        },
        updateHNO3(state,newHNO3){
            state.HNO3 = newHNO3;
        },
        updateHOO(state,newHOO){
            state.HOO = newHOO;
        },
        updateElectrical(state,newElectrical){
            state.Electrical = newElectrical;
        },
    },
    actions: {
        async login({commit},{companyid,password}){
            const res=await axios.post('http://47.97.176.174:8087/chainmaker?cmb=CompanyLogin',{
                CompanyID: companyid,
                Password: password
            })

            if(companyid == '' || password==''){
                ElMessage.error('请输入用户名和密码')
            }else {
                if (res.data.msg == 'success'){
                    const response=await axios.post('http://47.97.176.174:8087/chainmaker?cmb=GetCompanyByCompanyId',{
                        CompanyID:companyid
                    })
                    const companyData = response.data.data;
                    commit('setCompanyData', companyData);
                    router.push('/DashBoard')
                    ElMessage({
                        message: '登录成功',
                        type: 'success',
                    })
                } else {
                    ElMessage.error('用户名和密码不匹配')
                }
            }
        }
    }
});
Object.keys(ElementPlusIconsVue).forEach(key => {
    app.component(key, ElementPlusIconsVue[key]);
})

app.use(router)
app.config.globalProperties.$http=axios
app.use(ElementPlus)
app.use(store)
app.mount('#app')
