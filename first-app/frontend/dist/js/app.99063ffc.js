(function(){"use strict";var t={973:function(t,n,e){var r=e(9242),o=e(3396);const i={id:"app",class:"container"},u={class:"row"},a={class:"col-md-6 offset-md-3 py-5"},s=(0,o._)("h1",null,"Generate a thumbnail of a website",-1),l={class:"form-group"},c=(0,o._)("div",{class:"form-group"},[(0,o._)("button",{class:"btn btn-primary"},"Generate!")],-1),f=["src"];function h(t,n,e,h,p,d){return(0,o.wg)(),(0,o.iD)("div",i,[(0,o._)("div",u,[(0,o._)("div",a,[s,(0,o._)("form",{onSubmit:n[1]||(n[1]=(0,r.iM)(((...t)=>d.makeWebsiteThumbnail&&d.makeWebsiteThumbnail(...t)),["prevent"]))},[(0,o._)("div",l,[(0,o.wy)((0,o._)("input",{"onUpdate:modelValue":n[0]||(n[0]=t=>p.websiteUrl=t),placeholder:"Enter a website",type:"text",id:"website-input",class:"form-control"},null,512),[[r.nr,p.websiteUrl]])]),c,(0,o._)("img",{src:p.thumbnailUrl},null,8,f)],32)])])])}var p=e(4311),d={name:"App",data(){return{websiteUrl:"",thumbnailUrl:""}},methods:{makeWebsiteThumbnail(){p.Z.post("https://shot.screenshotapi.net/screenshot?token=QX33FDK-34944RG-PVSFV74-2PVDP0X&url=https%3A%2F%2Fgoogle.com&output=image&file_type=png&wait_for_event=load",{token:"QX33FDK-34944RG-PVSFV74-2PVDP0X",url:this.websiteUrl,width:1920,height:1080,output:"json",thumbnail_width:300}).then((t=>{this.thumbnailUrl=t.data.screenshot})).catch((t=>{window.alert(`The API returned an error: ${t}`)}))}}},b=e(89);const m=(0,b.Z)(d,[["render",h]]);var v=m;(0,r.ri)(v).mount("#app")}},n={};function e(r){var o=n[r];if(void 0!==o)return o.exports;var i=n[r]={exports:{}};return t[r](i,i.exports,e),i.exports}e.m=t,function(){var t=[];e.O=function(n,r,o,i){if(!r){var u=1/0;for(c=0;c<t.length;c++){r=t[c][0],o=t[c][1],i=t[c][2];for(var a=!0,s=0;s<r.length;s++)(!1&i||u>=i)&&Object.keys(e.O).every((function(t){return e.O[t](r[s])}))?r.splice(s--,1):(a=!1,i<u&&(u=i));if(a){t.splice(c--,1);var l=o();void 0!==l&&(n=l)}}return n}i=i||0;for(var c=t.length;c>0&&t[c-1][2]>i;c--)t[c]=t[c-1];t[c]=[r,o,i]}}(),function(){e.n=function(t){var n=t&&t.__esModule?function(){return t["default"]}:function(){return t};return e.d(n,{a:n}),n}}(),function(){e.d=function(t,n){for(var r in n)e.o(n,r)&&!e.o(t,r)&&Object.defineProperty(t,r,{enumerable:!0,get:n[r]})}}(),function(){e.g=function(){if("object"===typeof globalThis)return globalThis;try{return this||new Function("return this")()}catch(t){if("object"===typeof window)return window}}()}(),function(){e.o=function(t,n){return Object.prototype.hasOwnProperty.call(t,n)}}(),function(){var t={143:0};e.O.j=function(n){return 0===t[n]};var n=function(n,r){var o,i,u=r[0],a=r[1],s=r[2],l=0;if(u.some((function(n){return 0!==t[n]}))){for(o in a)e.o(a,o)&&(e.m[o]=a[o]);if(s)var c=s(e)}for(n&&n(r);l<u.length;l++)i=u[l],e.o(t,i)&&t[i]&&t[i][0](),t[i]=0;return e.O(c)},r=self["webpackChunkfrontend"]=self["webpackChunkfrontend"]||[];r.forEach(n.bind(null,0)),r.push=n.bind(null,r.push.bind(r))}();var r=e.O(void 0,[998],(function(){return e(973)}));r=e.O(r)})();
//# sourceMappingURL=app.99063ffc.js.map