"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[433],{87784:function(F,M,i){i.d(M,{Z:function(){return A}});var r=i(87462),w=i(67294),y={icon:{tag:"svg",attrs:{viewBox:"64 64 896 896",focusable:"false"},children:[{tag:"path",attrs:{d:"M512 64C264.6 64 64 264.6 64 512s200.6 448 448 448 448-200.6 448-448S759.4 64 512 64zm0 820c-205.4 0-372-166.6-372-372 0-89 31.3-170.8 83.5-234.8l523.3 523.3C682.8 852.7 601 884 512 884zm288.5-137.2L277.2 223.5C341.2 171.3 423 140 512 140c205.4 0 372 166.6 372 372 0 89-31.3 170.8-83.5 234.8z"}}]},name:"stop",theme:"outlined"},I=y,z=i(84089),B=function(H,N){return w.createElement(z.Z,(0,r.Z)({},H,{ref:N,icon:I}))},L=w.forwardRef(B),A=L},26412:function(F,M,i){i.d(M,{Z:function(){return ie}});var r=i(67294),w=i(93967),y=i.n(w),I=i(74443),z=i(53124),B=i(98675),L=i(25378),W={xxl:3,xl:3,lg:3,md:3,sm:2,xs:1},N=r.createContext({}),V=i(50344),J=function(e,t){var o={};for(var n in e)Object.prototype.hasOwnProperty.call(e,n)&&t.indexOf(n)<0&&(o[n]=e[n]);if(e!=null&&typeof Object.getOwnPropertySymbols=="function")for(var l=0,n=Object.getOwnPropertySymbols(e);l<n.length;l++)t.indexOf(n[l])<0&&Object.prototype.propertyIsEnumerable.call(e,n[l])&&(o[n[l]]=e[n[l]]);return o};const K=e=>(0,V.Z)(e).map(t=>Object.assign(Object.assign({},t==null?void 0:t.props),{key:t.key}));function Q(e,t,o){const n=r.useMemo(()=>t||K(o),[t,o]);return r.useMemo(()=>n.map(a=>{var{span:s}=a,c=J(a,["span"]);return s==="filled"?Object.assign(Object.assign({},c),{filled:!0}):Object.assign(Object.assign({},c),{span:typeof s=="number"?s:(0,I.m9)(e,s)})}),[n,e])}var Y=function(e,t){var o={};for(var n in e)Object.prototype.hasOwnProperty.call(e,n)&&t.indexOf(n)<0&&(o[n]=e[n]);if(e!=null&&typeof Object.getOwnPropertySymbols=="function")for(var l=0,n=Object.getOwnPropertySymbols(e);l<n.length;l++)t.indexOf(n[l])<0&&Object.prototype.propertyIsEnumerable.call(e,n[l])&&(o[n[l]]=e[n[l]]);return o};function q(e,t){let o=[],n=[],l=!1,a=0;return e.filter(s=>s).forEach(s=>{const{filled:c}=s,d=Y(s,["filled"]);if(c){n.push(d),o.push(n),n=[],a=0;return}const m=t-a;a+=s.span||1,a>=t?(a>t?(l=!0,n.push(Object.assign(Object.assign({},d),{span:m}))):n.push(d),o.push(n),n=[],a=0):n.push(d)}),n.length>0&&o.push(n),o=o.map(s=>{const c=s.reduce((d,m)=>d+(m.span||1),0);if(c<t){const d=s[s.length-1];return d.span=t-c+1,s}return s}),[o,l]}var k=(e,t)=>{const[o,n]=(0,r.useMemo)(()=>q(t,e),[t,e]);return o},_=e=>{let{children:t}=e;return t};function Z(e){return e!=null}var R=e=>{const{itemPrefixCls:t,component:o,span:n,className:l,style:a,labelStyle:s,contentStyle:c,bordered:d,label:m,content:u,colon:v,type:$}=e,b=o;return d?r.createElement(b,{className:y()({[`${t}-item-label`]:$==="label",[`${t}-item-content`]:$==="content"},l),style:a,colSpan:n},Z(m)&&r.createElement("span",{style:s},m),Z(u)&&r.createElement("span",{style:c},u)):r.createElement(b,{className:y()(`${t}-item`,l),style:a,colSpan:n},r.createElement("div",{className:`${t}-item-container`},(m||m===0)&&r.createElement("span",{className:y()(`${t}-item-label`,{[`${t}-item-no-colon`]:!v}),style:s},m),(u||u===0)&&r.createElement("span",{className:y()(`${t}-item-content`),style:c},u)))};function T(e,t,o){let{colon:n,prefixCls:l,bordered:a}=t,{component:s,type:c,showLabel:d,showContent:m,labelStyle:u,contentStyle:v}=o;return e.map(($,b)=>{let{label:O,children:P,prefixCls:S=l,className:x,style:C,labelStyle:g,contentStyle:p,span:h=1,key:j}=$;return typeof s=="string"?r.createElement(R,{key:`${c}-${j||b}`,className:x,style:C,labelStyle:Object.assign(Object.assign({},u),g),contentStyle:Object.assign(Object.assign({},v),p),span:h,colon:n,component:s,itemPrefixCls:S,bordered:a,label:d?O:null,content:m?P:null,type:c}):[r.createElement(R,{key:`label-${j||b}`,className:x,style:Object.assign(Object.assign(Object.assign({},u),C),g),span:1,colon:n,component:s[0],itemPrefixCls:S,bordered:a,label:O,type:"label"}),r.createElement(R,{key:`content-${j||b}`,className:x,style:Object.assign(Object.assign(Object.assign({},v),C),p),span:h*2-1,component:s[1],itemPrefixCls:S,bordered:a,content:P,type:"content"})]})}var ee=e=>{const t=r.useContext(N),{prefixCls:o,vertical:n,row:l,index:a,bordered:s}=e;return n?r.createElement(r.Fragment,null,r.createElement("tr",{key:`label-${a}`,className:`${o}-row`},T(l,e,Object.assign({component:"th",type:"label",showLabel:!0},t))),r.createElement("tr",{key:`content-${a}`,className:`${o}-row`},T(l,e,Object.assign({component:"td",type:"content",showContent:!0},t)))):r.createElement("tr",{key:a,className:`${o}-row`},T(l,e,Object.assign({component:s?["th","td"]:"td",type:"item",showLabel:!0,showContent:!0},t)))},f=i(11568),X=i(14747),te=i(83559),ne=i(83262);const le=e=>{const{componentCls:t,labelBg:o}=e;return{[`&${t}-bordered`]:{[`> ${t}-view`]:{border:`${(0,f.bf)(e.lineWidth)} ${e.lineType} ${e.colorSplit}`,"> table":{tableLayout:"auto"},[`${t}-row`]:{borderBottom:`${(0,f.bf)(e.lineWidth)} ${e.lineType} ${e.colorSplit}`,"&:last-child":{borderBottom:"none"},[`> ${t}-item-label, > ${t}-item-content`]:{padding:`${(0,f.bf)(e.padding)} ${(0,f.bf)(e.paddingLG)}`,borderInlineEnd:`${(0,f.bf)(e.lineWidth)} ${e.lineType} ${e.colorSplit}`,"&:last-child":{borderInlineEnd:"none"}},[`> ${t}-item-label`]:{color:e.colorTextSecondary,backgroundColor:o,"&::after":{display:"none"}}}},[`&${t}-middle`]:{[`${t}-row`]:{[`> ${t}-item-label, > ${t}-item-content`]:{padding:`${(0,f.bf)(e.paddingSM)} ${(0,f.bf)(e.paddingLG)}`}}},[`&${t}-small`]:{[`${t}-row`]:{[`> ${t}-item-label, > ${t}-item-content`]:{padding:`${(0,f.bf)(e.paddingXS)} ${(0,f.bf)(e.padding)}`}}}}}},oe=e=>{const{componentCls:t,extraColor:o,itemPaddingBottom:n,itemPaddingEnd:l,colonMarginRight:a,colonMarginLeft:s,titleMarginBottom:c}=e;return{[t]:Object.assign(Object.assign(Object.assign({},(0,X.Wf)(e)),le(e)),{"&-rtl":{direction:"rtl"},[`${t}-header`]:{display:"flex",alignItems:"center",marginBottom:c},[`${t}-title`]:Object.assign(Object.assign({},X.vS),{flex:"auto",color:e.titleColor,fontWeight:e.fontWeightStrong,fontSize:e.fontSizeLG,lineHeight:e.lineHeightLG}),[`${t}-extra`]:{marginInlineStart:"auto",color:o,fontSize:e.fontSize},[`${t}-view`]:{width:"100%",borderRadius:e.borderRadiusLG,table:{width:"100%",tableLayout:"fixed",borderCollapse:"collapse"}},[`${t}-row`]:{"> th, > td":{paddingBottom:n,paddingInlineEnd:l},"> th:last-child, > td:last-child":{paddingInlineEnd:0},"&:last-child":{borderBottom:"none","> th, > td":{paddingBottom:0}}},[`${t}-item-label`]:{color:e.colorTextTertiary,fontWeight:"normal",fontSize:e.fontSize,lineHeight:e.lineHeight,textAlign:"start","&::after":{content:'":"',position:"relative",top:-.5,marginInline:`${(0,f.bf)(s)} ${(0,f.bf)(a)}`},[`&${t}-item-no-colon::after`]:{content:'""'}},[`${t}-item-no-label`]:{"&::after":{margin:0,content:'""'}},[`${t}-item-content`]:{display:"table-cell",flex:1,color:e.contentColor,fontSize:e.fontSize,lineHeight:e.lineHeight,wordBreak:"break-word",overflowWrap:"break-word"},[`${t}-item`]:{paddingBottom:0,verticalAlign:"top","&-container":{display:"flex",[`${t}-item-label`]:{display:"inline-flex",alignItems:"baseline"},[`${t}-item-content`]:{display:"inline-flex",alignItems:"baseline",minWidth:"1em"}}},"&-middle":{[`${t}-row`]:{"> th, > td":{paddingBottom:e.paddingSM}}},"&-small":{[`${t}-row`]:{"> th, > td":{paddingBottom:e.paddingXS}}}})}},re=e=>({labelBg:e.colorFillAlter,titleColor:e.colorText,titleMarginBottom:e.fontSizeSM*e.lineHeightSM,itemPaddingBottom:e.padding,itemPaddingEnd:e.padding,colonMarginRight:e.marginXS,colonMarginLeft:e.marginXXS/2,contentColor:e.colorText,extraColor:e.colorText});var se=(0,te.I$)("Descriptions",e=>{const t=(0,ne.IX)(e,{});return oe(t)},re),ae=function(e,t){var o={};for(var n in e)Object.prototype.hasOwnProperty.call(e,n)&&t.indexOf(n)<0&&(o[n]=e[n]);if(e!=null&&typeof Object.getOwnPropertySymbols=="function")for(var l=0,n=Object.getOwnPropertySymbols(e);l<n.length;l++)t.indexOf(n[l])<0&&Object.prototype.propertyIsEnumerable.call(e,n[l])&&(o[n[l]]=e[n[l]]);return o};const G=e=>{const{prefixCls:t,title:o,extra:n,column:l,colon:a=!0,bordered:s,layout:c,children:d,className:m,rootClassName:u,style:v,size:$,labelStyle:b,contentStyle:O,items:P}=e,S=ae(e,["prefixCls","title","extra","column","colon","bordered","layout","children","className","rootClassName","style","size","labelStyle","contentStyle","items"]),{getPrefixCls:x,direction:C,descriptions:g}=r.useContext(z.E_),p=x("descriptions",t),h=(0,L.Z)(),j=r.useMemo(()=>{var E;return typeof l=="number"?l:(E=(0,I.m9)(h,Object.assign(Object.assign({},W),l)))!==null&&E!==void 0?E:3},[h,l]),ce=Q(h,P,d),D=(0,B.Z)($),de=k(j,ce),[me,pe,fe]=se(p),ue=r.useMemo(()=>({labelStyle:b,contentStyle:O}),[b,O]);return me(r.createElement(N.Provider,{value:ue},r.createElement("div",Object.assign({className:y()(p,g==null?void 0:g.className,{[`${p}-${D}`]:D&&D!=="default",[`${p}-bordered`]:!!s,[`${p}-rtl`]:C==="rtl"},m,u,pe,fe),style:Object.assign(Object.assign({},g==null?void 0:g.style),v)},S),(o||n)&&r.createElement("div",{className:`${p}-header`},o&&r.createElement("div",{className:`${p}-title`},o),n&&r.createElement("div",{className:`${p}-extra`},n)),r.createElement("div",{className:`${p}-view`},r.createElement("table",null,r.createElement("tbody",null,de.map((E,U)=>r.createElement(ee,{key:U,index:U,colon:a,prefixCls:p,vertical:c==="vertical",bordered:s,row:E}))))))))};G.Item=_;var ie=G}}]);