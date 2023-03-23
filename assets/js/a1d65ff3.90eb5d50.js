"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[2768],{3905:(e,t,n)=>{n.d(t,{Zo:()=>s,kt:()=>c});var o=n(7294);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);t&&(o=o.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,o)}return n}function r(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function p(e,t){if(null==e)return{};var n,o,i=function(e,t){if(null==e)return{};var n,o,i={},a=Object.keys(e);for(o=0;o<a.length;o++)n=a[o],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(o=0;o<a.length;o++)n=a[o],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var l=o.createContext({}),d=function(e){var t=o.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):r(r({},t),e)),n},s=function(e){var t=d(e.components);return o.createElement(l.Provider,{value:t},e.children)},m={inlineCode:"code",wrapper:function(e){var t=e.children;return o.createElement(o.Fragment,{},t)}},u=o.forwardRef((function(e,t){var n=e.components,i=e.mdxType,a=e.originalType,l=e.parentName,s=p(e,["components","mdxType","originalType","parentName"]),u=d(n),c=i,g=u["".concat(l,".").concat(c)]||u[c]||m[c]||a;return n?o.createElement(g,r(r({ref:t},s),{},{components:n})):o.createElement(g,r({ref:t},s))}));function c(e,t){var n=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var a=n.length,r=new Array(a);r[0]=u;var p={};for(var l in t)hasOwnProperty.call(t,l)&&(p[l]=t[l]);p.originalType=e,p.mdxType="string"==typeof e?e:i,r[1]=p;for(var d=2;d<a;d++)r[d]=n[d];return o.createElement.apply(null,r)}return o.createElement.apply(null,n)}u.displayName="MDXCreateElement"},5283:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>r,default:()=>m,frontMatter:()=>a,metadata:()=>p,toc:()=>d});var o=n(7462),i=(n(7294),n(3905));const a={sidebar_position:1},r="Dependency Injection",p={unversionedId:"building-modules/depinject",id:"version-v0.47/building-modules/depinject",title:"Dependency Injection",description:"Pre-requisite Readings",source:"@site/versioned_docs/version-v0.47/building-modules/15-depinject.md",sourceDirName:"building-modules",slug:"/building-modules/depinject",permalink:"/v0.47/building-modules/depinject",draft:!1,tags:[],version:"v0.47",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"Module Simulation",permalink:"/v0.47/building-modules/simulator"},next:{title:"Testing",permalink:"/v0.47/building-modules/testing"}},l={},d=[{value:"Module Configuration",id:"module-configuration",level:2},{value:"Dependency Definition",id:"dependency-definition",level:2},{value:"App Wiring",id:"app-wiring",level:2}],s={toc:d};function m(e){let{components:t,...n}=e;return(0,i.kt)("wrapper",(0,o.Z)({},s,n,{components:t,mdxType:"MDXLayout"}),(0,i.kt)("h1",{id:"dependency-injection"},"Dependency Injection"),(0,i.kt)("admonition",{type:"note"},(0,i.kt)("h3",{parentName:"admonition",id:"pre-requisite-readings"},"Pre-requisite Readings"),(0,i.kt)("ul",{parentName:"admonition"},(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("a",{parentName:"li",href:"/v0.47/tooling/depinject"},"Cosmos SDK Dependency Injection Framework")))),(0,i.kt)("p",null,(0,i.kt)("a",{parentName:"p",href:"/v0.47/tooling/depinject"},(0,i.kt)("inlineCode",{parentName:"a"},"depinject"))," is used to wire any module in ",(0,i.kt)("inlineCode",{parentName:"p"},"app.go"),".\nAll core modules are already configured to support dependency injection."),(0,i.kt)("p",null,"To work with ",(0,i.kt)("inlineCode",{parentName:"p"},"depinject")," a module must define its configuration and requirements so that ",(0,i.kt)("inlineCode",{parentName:"p"},"depinject")," can provide the right dependencies."),(0,i.kt)("p",null,"In brief, as a module developer, the following steps are required:"),(0,i.kt)("ol",null,(0,i.kt)("li",{parentName:"ol"},"Define the module configuration using Protobuf"),(0,i.kt)("li",{parentName:"ol"},"Define the module dependencies in ",(0,i.kt)("inlineCode",{parentName:"li"},"x/{moduleName}/module.go"))),(0,i.kt)("p",null,"A chain developer can then use the module by following these two steps:"),(0,i.kt)("ol",null,(0,i.kt)("li",{parentName:"ol"},"Configure the module in ",(0,i.kt)("inlineCode",{parentName:"li"},"app_config.go")," or ",(0,i.kt)("inlineCode",{parentName:"li"},"app.yaml")),(0,i.kt)("li",{parentName:"ol"},"Inject the module in ",(0,i.kt)("inlineCode",{parentName:"li"},"app.go"))),(0,i.kt)("h2",{id:"module-configuration"},"Module Configuration"),(0,i.kt)("p",null,"The module available configuration is defined in a Protobuf file, located at ",(0,i.kt)("inlineCode",{parentName:"p"},"{moduleName}/module/v1/module.proto"),"."),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-protobuf",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/proto/cosmos/group/module/v1/module.proto\n")),(0,i.kt)("ul",null,(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},(0,i.kt)("inlineCode",{parentName:"p"},"go_import")," must point to the Go package of the custom module.")),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("p",{parentName:"li"},"Message fields define the module configuration.\nThat configuration can be set in the ",(0,i.kt)("inlineCode",{parentName:"p"},"app_config.go")," / ",(0,i.kt)("inlineCode",{parentName:"p"},"app.yaml")," file for a chain developer to configure the module.",(0,i.kt)("br",{parentName:"p"}),"\n","Taking ",(0,i.kt)("inlineCode",{parentName:"p"},"group")," as example, a chain developer is able to decide, thanks to ",(0,i.kt)("inlineCode",{parentName:"p"},"uint64 max_metadata_len"),", what the maximum metatada length allowed for a group porposal is."),(0,i.kt)("pre",{parentName:"li"},(0,i.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/simapp/app_config.go#L226-L230\n")))),(0,i.kt)("p",null,"That message is generated using ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/scripts/protocgen-pulsar.sh"},(0,i.kt)("inlineCode",{parentName:"a"},"pulsar"))," (by running ",(0,i.kt)("inlineCode",{parentName:"p"},"make proto-gen"),").\nIn the case of the ",(0,i.kt)("inlineCode",{parentName:"p"},"group")," module, this file is generated here: ",(0,i.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/api/cosmos/group/module/v1/module.pulsar.go"},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/api/cosmos/group/module/v1/module.pulsar.go"),"."),(0,i.kt)("p",null,"The part that is relevant for the module configuration is:"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/api/cosmos/group/module/v1/module.pulsar.go#L515-L527\n")),(0,i.kt)("admonition",{type:"note"},(0,i.kt)("p",{parentName:"admonition"},"Pulsar is optional. The official ",(0,i.kt)("a",{parentName:"p",href:"https://developers.google.com/protocol-buffers/docs/reference/go-generated"},(0,i.kt)("inlineCode",{parentName:"a"},"protoc-gen-go"))," can be used as well.")),(0,i.kt)("h2",{id:"dependency-definition"},"Dependency Definition"),(0,i.kt)("p",null,"Once the configuration proto is defined, the module's ",(0,i.kt)("inlineCode",{parentName:"p"},"module.go")," must define what dependencies are required by the module.\nThe boilerplate is similar for all modules."),(0,i.kt)("admonition",{type:"warning"},(0,i.kt)("p",{parentName:"admonition"},"All methods, structs and their fields must be public for ",(0,i.kt)("inlineCode",{parentName:"p"},"depinject"),".")),(0,i.kt)("ol",null,(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("p",{parentName:"li"},"Import the module configuration generated package:"),(0,i.kt)("pre",{parentName:"li"},(0,i.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/group/module/module.go#L12-L14\n")),(0,i.kt)("p",{parentName:"li"},"Define an ",(0,i.kt)("inlineCode",{parentName:"p"},"init()")," function for defining the ",(0,i.kt)("inlineCode",{parentName:"p"},"providers")," of the module configuration:",(0,i.kt)("br",{parentName:"p"}),"\n","This registers the module configuration message and the wiring of the module."),(0,i.kt)("pre",{parentName:"li"},(0,i.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/group/module/module.go#L199-L204\n"))),(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("p",{parentName:"li"},"Define a struct that inherits ",(0,i.kt)("inlineCode",{parentName:"p"},"depinject.In")," and define the module inputs (i.e. module dependencies):"),(0,i.kt)("ul",{parentName:"li"},(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"depinject")," provides the right dependencies to the module."),(0,i.kt)("li",{parentName:"ul"},(0,i.kt)("inlineCode",{parentName:"li"},"depinject")," also checks that all dependencies are provided.")),(0,i.kt)("admonition",{parentName:"li",type:"tip"},(0,i.kt)("p",{parentName:"admonition"},"For making a dependency optional, add the ",(0,i.kt)("inlineCode",{parentName:"p"},'optional:"true"')," struct tag.  ")),(0,i.kt)("pre",{parentName:"li"},(0,i.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/group/module/module.go#L206-L216\n"))),(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("p",{parentName:"li"},"Define the module outputs with a public struct that inherits ",(0,i.kt)("inlineCode",{parentName:"p"},"depinject.Out"),":\nThe module outputs are the dependencies that the module provides to other modules. It is usually the module itself and its keeper."),(0,i.kt)("pre",{parentName:"li"},(0,i.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/group/module/module.go#L218-L223\n"))),(0,i.kt)("li",{parentName:"ol"},(0,i.kt)("p",{parentName:"li"},"Create a function named ",(0,i.kt)("inlineCode",{parentName:"p"},"ProvideModule")," (as called in 1.) and use the inputs for instantiating the module outputs."),(0,i.kt)("pre",{parentName:"li"},(0,i.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/group/module/module.go#L225-L235\n")))),(0,i.kt)("p",null,"The ",(0,i.kt)("inlineCode",{parentName:"p"},"ProvideModule")," function should return an instance of ",(0,i.kt)("inlineCode",{parentName:"p"},"cosmossdk.io/core/appmodule.AppModule")," which implements\none or more app module extension interfaces for initializing the module."),(0,i.kt)("p",null,"Following is the complete app wiring configuration for ",(0,i.kt)("inlineCode",{parentName:"p"},"group"),":"),(0,i.kt)("pre",null,(0,i.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/group/module/module.go#L195-L235\n")),(0,i.kt)("p",null,"The module is now ready to be used with ",(0,i.kt)("inlineCode",{parentName:"p"},"depinject")," by a chain developer."),(0,i.kt)("h2",{id:"app-wiring"},"App Wiring"),(0,i.kt)("p",null,"The App Wiring is done in ",(0,i.kt)("inlineCode",{parentName:"p"},"app_config.go")," / ",(0,i.kt)("inlineCode",{parentName:"p"},"app.yaml")," and ",(0,i.kt)("inlineCode",{parentName:"p"},"app_v2.go")," and is explained in detail in the ",(0,i.kt)("a",{parentName:"p",href:"/v0.47/building-apps/app-go-v2"},"overview of ",(0,i.kt)("inlineCode",{parentName:"a"},"app_v2.go")),"."))}m.isMDXComponent=!0}}]);