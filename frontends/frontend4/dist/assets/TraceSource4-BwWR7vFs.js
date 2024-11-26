import{_ as V,j as B,u as T,r,o as q,a as d,b as A,c as z,d as D,e,w as a,f as o,t as c,g as J,p as $,h as j}from"./index-Dqiu3Z7p.js";const w=i=>($("data-v-4734bf89"),i=i(),j(),i),M={class:"traceSource4"},H={style:{display:"flex","flex-direction":"row","margin-top":"30px"}},K=w(()=>D("p",{style:{"text-align":"left","margin-left":"20px","font-size":"18px","font-weight":"bold"}},"六氟化硫回收",-1)),Q=w(()=>D("p",{style:{"text-align":"left","margin-top":"50px","margin-left":"20px","font-size":"18px","font-weight":"bold"}},"输配电损失",-1)),U={__name:"TraceSource4",setup(i){const C=B(),h=T(),P=J(),{proxy:k}=P,I=C.query.Id,G=C.query.CompanyId,l=r({}),u=r({}),_=r([]),m=r([]),f=r(0),b=r(0),v=r(0),x=r(0),g=r(0),y=r(0);q(async()=>{l.value.Id=I,console.log(I),l.value.CompanyId=G,await k.$axios.post("http://47.97.176.174:8087/chainmaker?cmb=SearchPGEById",{Id:parseInt(l.value.Id)}).then(t=>{console.log(t),u.value=t.data.data,_.value=[...t.data.data.Data.PGERepairREC],m.value=[...t.data.data.Data.PGERetireREC],f.value=t.data.data.Data.ELSW,b.value=t.data.data.Data.ELSR,v.value=t.data.data.Data.ELSC,x.value=t.data.data.Data.ELSD,g.value=t.data.data.Data.ELSPDL,y.value=t.data.data.Data.EFDW,l.value.currentCompanyData=u,l.value.PGERepairREC=_,l.value.PGERetireREC=m,l.value.ELSW=f,l.value.ELSR=b,l.value.ELSC=v,l.value.ELSD=x,l.value.ELSPDL=g,l.value.EFDW=y}).catch(function(t){console.log(t)});const s=F();s&&(l.value=s)});const O=s=>{console.log(s),sessionStorage.setItem("myData4",JSON.stringify(s)),h.push("/traceSource3/"+s.CompanyId)},W=()=>{sessionStorage.removeItem("myData1"),sessionStorage.removeItem("myData2"),sessionStorage.removeItem("myData3"),sessionStorage.removeItem("myData4"),h.push("/carbonAccountingData")},F=()=>{const s=sessionStorage.getItem("myData4");return s?JSON.parse(s):null};return(s,t)=>{const R=d("el-button"),n=d("el-descriptions-item"),E=d("el-descriptions"),S=d("el-tab-pane"),p=d("el-table-column"),L=d("el-table"),N=d("el-tabs");return A(),z("div",M,[D("div",H,[e(R,{type:"primary",icon:"ArrowLeft",onClick:t[0]||(t[0]=X=>O(l.value)),style:{width:"120px","margin-left":"20px"}},{default:a(()=>[o("返回")]),_:1}),e(R,{type:"default",icon:"DArrowLeft",onClick:W,style:{width:"120px","margin-left":"20px"}},{default:a(()=>[o("退出")]),_:1})]),e(N,{type:"border-card",style:{margin:"30px 20px 30px 20px",height:"900px"}},{default:a(()=>[e(S,{label:"附表1：报告主体 2024 年二氧化碳排放量报告"},{default:a(()=>[e(E,{column:1,border:"",style:{margin:"20px 10px 20px 10px"}},{default:a(()=>[e(n,{label:"使用六氟化硫设备修理与退役过程产生的排放(tCO2)",width:"200"},{default:a(()=>[o(c(u.value.FSCO),1)]),_:1}),e(n,{label:"输配电引起的二氧化碳排放(tCO2)"},{default:a(()=>[o(c(u.value.SPDCO),1)]),_:1}),e(n,{label:"企业二氧化碳排放总量(tCO2)"},{default:a(()=>[o(c(u.value.Total),1)]),_:1})]),_:1})]),_:1}),e(S,{label:"附表2：报告主体活动水平数据"},{default:a(()=>[K,e(L,{data:_.value,border:"","header-cell-style":{background:"#f4f4f5"},style:{width:"96%","margin-left":"15px","margin-top":"20px"}},{default:a(()=>[e(p,{label:"修理设备",prop:"Id"}),e(p,{label:"设备容量(千克)",prop:"RECV"}),e(p,{label:"实际回收量(千克)",prop:"RECR"})]),_:1},8,["data"]),e(L,{data:m.value,border:"","header-cell-style":{background:"#f4f4f5"},style:{width:"96%","margin-left":"15px","margin-top":"20px"}},{default:a(()=>[e(p,{label:"退役设备",prop:"Id"}),e(p,{label:"设备容量(千克)",prop:"RECV"}),e(p,{label:"实际回收量(千克)",prop:"RECR"})]),_:1},8,["data"]),Q,e(E,{column:1,border:"",style:{margin:"20px 10px 20px 10px"}},{default:a(()=>[e(n,{label:"电厂上网电量(兆瓦时)",width:"200"},{default:a(()=>[o(c(f.value),1)]),_:1}),e(n,{label:"自外省输入电量(兆瓦时)"},{default:a(()=>[o(c(b.value),1)]),_:1}),e(n,{label:"向外省输出电量(兆瓦时)"},{default:a(()=>[o(c(v.value),1)]),_:1}),e(n,{label:"售电量(兆瓦时)"},{default:a(()=>[o(c(x.value),1)]),_:1}),e(n,{label:"输配电量(兆瓦时)"},{default:a(()=>[o(c(g.value),1)]),_:1})]),_:1})]),_:1}),e(S,{label:"附表3：报告主体排放因子"},{default:a(()=>[e(E,{column:3,border:"",direction:"vertical",style:{margin:"20px 10px 20px 10px"}},{default:a(()=>[e(n,{label:"名称",width:"200"},{default:a(()=>[o("电力")]),_:1}),e(n,{label:"排放因子单位"},{default:a(()=>[o("吨二氧化碳/兆瓦时")]),_:1}),e(n,{label:"二氧化碳排放因子"},{default:a(()=>[o(c(y.value),1)]),_:1})]),_:1})]),_:1})]),_:1})])}}},Z=V(U,[["__scopeId","data-v-4734bf89"]]);export{Z as default};