"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[1610],{80821:function(A,f,t){t.d(f,{Z:function(){return T}});var p=t(5574),a=t.n(p),c=t(24963),s=t(86738),l=t(28036),O=t(83062),E=t(67294),h=t(85893);function T(u){var R=(0,E.useState)(!1),P=a()(R,2),D=P[0],M=P[1],v=(0,E.useContext)(c.Z),y=v.lang,j=v.notice,C=v.message,_=function(){M(!0),u.action().then(function(g){if(M(!1),typeof u.onSuccess=="function"&&u.onSuccess(g),u.messageSuccess){var o="";typeof u.messageSuccess=="function"?o=u.messageSuccess(g):o=u.messageSuccess,o=y("notification.".concat(o))+y("notification.title.success"),u.asynced?C.info(o):C.success(o)}}).catch(function(g){M(!1),typeof u.onError=="function"&&u.onError(g)})};return u.confirm?(0,h.jsx)(s.Z,{style:{width:500},title:u.confirmTitle?u.confirmTitle:y("notification.title.tip"),description:u.confirm,onConfirm:_,okText:"Yes",cancelText:"No",children:(0,h.jsx)(l.ZP,{disabled:u.disabled,icon:u.icon,loading:D,danger:u.danger,type:u.type,children:u.children})}):(0,h.jsx)(O.Z,{title:u.tips,children:(0,h.jsx)(l.ZP,{disabled:u.disabled,icon:u.icon,loading:D,onClick:_,danger:u.danger,type:u.type,children:u.children})})}},78451:function(A,f,t){t.d(f,{Z:function(){return c}});var p=t(67294),a=t(85893);function c(s){return(0,a.jsx)("span",{style:{width:s.width?s.width:"auto",wordBreak:"break-word"},children:s.content})}},24963:function(A,f,t){var p=t(67294),a=(0,p.createContext)({});f.Z=a},67108:function(A,f,t){t.r(f),t.d(f,{default:function(){return U}});var p=t(15009),a=t.n(p),c=t(99289),s=t.n(c),l=t(90078),O=t(10641),E=t(62597),h=t(28036),T=t(83062),u=t(42075),R=t(50136),P=t(63713),D=t(43425),M=t(67294),v=t(96974),y=t(80821),j=t(67255),C=t(78451),_=t(85893);function U(){var g=(0,v.UO)(),o=(0,M.useRef)();return(0,M.useEffect)(function(){var d;(d=o.current)===null||d===void 0||d.reload()},[g]),(0,_.jsx)(l._z,{header:{extra:[(0,_.jsx)(h.ZP,{type:"primary",children:(0,_.jsx)(P.Link,{to:"/app/create/image",children:"\u521B\u5EFA\u5BB9\u5668"})},"create")]},children:(0,_.jsx)(O.Z,{scroll:{x:"max-content"},actionRef:o,columns:[{title:"\u540D\u79F0",dataIndex:"siteTitle"},{title:"\u6807\u8BC6",dataIndex:"siteName"},{title:"\u6700\u540E\u9519\u8BEF",dataIndex:"message",width:250,render:function(n,r,i,e,m){return(0,_.jsx)(C.Z,{content:r.message})}},{title:"\u5220\u9664\u65E5\u671F",ellipsis:!0,dataIndex:"deletedAt",search:!1,width:180,render:function(n,r,i,e,m){return new Date(r.deletedAt).toLocaleString()}},{title:"\u64CD\u4F5C",valueType:"option",key:"option",width:100,render:function(n,r,i,e){return[(0,_.jsx)(P.Link,{to:"/app/create/image?op=update&id="+r.id,type:"link",children:(0,_.jsx)(T.Z,{title:"\u518D\u6B21\u6784\u5EFA",children:(0,_.jsx)(D.Z,{})})},"rebuild")]}}],rowKey:"id",request:function(){var d=s()(a()().mark(function n(r,i,e){var m,L,B;return a()().wrap(function(b){for(;;)switch(b.prev=b.next){case 0:return b.next=2,(0,E.cl)({pageSize:(m=r.pageSize)!==null&&m!==void 0?m:10,page:(L=r.current)!==null&&L!==void 0?L:1,isDelete:!0,siteTitle:r.siteTitle,siteName:r.siteName});case 2:return B=b.sent,b.abrupt("return",{data:B.data.list,success:!0,total:B.data.total});case 4:case"end":return b.stop()}},n)}));return function(n,r,i){return d.apply(this,arguments)}}(),rowSelection:{defaultSelectedRowKeys:[],alwaysShowAlert:!0},tableAlertOptionRender:function(n){var r=n.selectedRowKeys;return(0,_.jsx)(u.Z,{size:16,children:(0,_.jsx)(y.Z,{danger:!0,type:"primary",action:function(){return(0,E.xb)({id:r})},onSuccess:function(){var e,m;return!((e=o.current)===null||e===void 0)&&e.reloadAndRest&&((m=o.current)===null||m===void 0||m.reloadAndRest()),!0},onError:function(){var e,m;return!((e=o.current)===null||e===void 0)&&e.reset&&((m=o.current)===null||m===void 0||m.reset()),!0},messageSuccess:"delete",children:"\u6279\u91CF\u5220\u9664"})})},pagination:(0,j.O)(),search:{collapsed:!1},tableExtraRender:function(){return(0,_.jsx)(R.Z,{mode:"horizontal",selectedKeys:["recycle"],items:[{label:(0,_.jsx)(P.Link,{to:"/app/list",replace:!0,children:"\u5BB9\u5668\u5217\u8868"}),key:"all"},{label:(0,_.jsx)(P.Link,{to:"/app/recycle",replace:!0,children:"\u56DE\u6536\u7AD9"}),key:"recycle"}]})}})})}},62597:function(A,f,t){t.d(f,{$G:function(){return O},Ct:function(){return y},Tb:function(){return u},XH:function(){return P},cl:function(){return h},iE:function(){return M},lK:function(){return U},xb:function(){return C},ze:function(){return o}});var p=t(15009),a=t.n(p),c=t(99289),s=t.n(c),l=t(63713);function O(n){return E.apply(this,arguments)}function E(){return E=s()(a()().mark(function n(r){return a()().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,l.request)("/api/app/site/create-by-image",{method:"POST",data:r}));case 1:case"end":return e.stop()}},n)})),E.apply(this,arguments)}function h(n){return T.apply(this,arguments)}function T(){return T=s()(a()().mark(function n(r){return a()().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,l.request)("/api/app/site/get-list",{method:"POST",data:r}));case 1:case"end":return e.stop()}},n)})),T.apply(this,arguments)}function u(n){return R.apply(this,arguments)}function R(){return R=s()(a()().mark(function n(r){return a()().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return r.download=!1,e.next=3,(0,l.request)("/api/app/log/run",{method:"POST",data:r});case 3:return e.abrupt("return",e.sent);case 4:case"end":return e.stop()}},n)})),R.apply(this,arguments)}function P(n){return D.apply(this,arguments)}function D(){return D=s()(a()().mark(function n(r){return a()().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return r.download=!0,e.next=3,(0,l.request)("/api/app/log/run",{method:"POST",data:r,responseType:"blob"});case 3:return e.abrupt("return",e.sent);case 4:case"end":return e.stop()}},n)})),D.apply(this,arguments)}function M(n){return v.apply(this,arguments)}function v(){return v=s()(a()().mark(function n(r){return a()().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,(0,l.request)("/api/app/site/get-detail",{data:r,method:"POST"});case 2:return e.abrupt("return",e.sent);case 3:case"end":return e.stop()}},n)})),v.apply(this,arguments)}function y(n){return j.apply(this,arguments)}function j(){return j=s()(a()().mark(function n(r){return a()().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,l.request)("/api/app/container/delete",{method:"POST",data:r}));case 1:case"end":return e.stop()}},n)})),j.apply(this,arguments)}function C(n){return _.apply(this,arguments)}function _(){return _=s()(a()().mark(function n(r){return a()().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,l.request)("/api/app/site/delete",{method:"POST",data:r}));case 1:case"end":return e.stop()}},n)})),_.apply(this,arguments)}function U(n){return g.apply(this,arguments)}function g(){return g=s()(a()().mark(function n(r){return a()().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,(0,l.request)("/api/app/site/update-title",{method:"POST",data:r});case 2:return e.abrupt("return",e.sent);case 3:case"end":return e.stop()}},n)})),g.apply(this,arguments)}function o(n){return d.apply(this,arguments)}function d(){return d=s()(a()().mark(function n(r){return a()().wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.abrupt("return",(0,l.request)("/api/app/container/ignore",{method:"POST",data:r}));case 1:case"end":return e.stop()}},n)})),d.apply(this,arguments)}},67255:function(A,f,t){t.d(f,{O:function(){return p},j:function(){return a}});function p(){var c,s=parseInt((c=localStorage.getItem("dpanel-pagesize"))!==null&&c!==void 0?c:"0");return{showSizeChanger:!0,defaultPageSize:s!=null?s:20}}function a(c){var s="dpanel-table-column-".concat(c),l={};if(localStorage.getItem(s)){var O;l=JSON.parse((O=localStorage.getItem(s))!==null&&O!==void 0?O:"{}")}return{defaultValue:l,onChange:function(h){localStorage.setItem("dpanel-table-column-".concat(c),JSON.stringify(h))}}}}}]);