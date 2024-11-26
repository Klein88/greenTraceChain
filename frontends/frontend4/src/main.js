// import './assets/main.css'
// import './assets/base.css'

import { createApp } from 'vue'
import App from './App.vue'
import router from './router'

import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import PrimeVue from "primevue/config";
import 'primevue/resources/themes/saga-blue/theme.css'
import 'primevue/resources/primevue.min.css'
import 'primeicons/primeicons.css'
import 'primeflex/primeflex.css'

import axios from 'axios'

import Avatar from 'primevue/avatar'
import Button from 'primevue/button'
import Card from 'primevue/card'
import ColorPicker from 'primevue/colorpicker'
import Column from 'primevue/column'
import DataTable from 'primevue/datatable'
import Panel from 'primevue/panel'
import Fieldset from 'primevue/fieldset'

const app = createApp(App)

app.use(router)

app.use(ElementPlus)

app.use(PrimeVue)

app.component('Avatar', Avatar)
app.component('Button', Button)
app.component('Card', Card)
app.component('ColorPicker', ColorPicker)
app.component('Column', Column)
app.component('DataTable', DataTable)
app.component('Panel', Panel)
app.component('Fieldset', Fieldset)

Object.keys(ElementPlusIconsVue).forEach(key => {
    app.component(key, ElementPlusIconsVue[key]);
})

app.config.globalProperties.$axios=axios
// app.provide('$axios', axios)

app.mount('#app')
