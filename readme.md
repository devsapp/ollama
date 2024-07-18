
> 注：当前项目为 Serverless Devs 应用，由于应用中会存在需要初始化才可运行的变量（例如应用部署地区、函数名等等），所以**不推荐**直接 Clone 本仓库到本地进行部署或直接复制 s.yaml 使用，**强烈推荐**通过 `s init ${模版名称}` 的方法或应用中心进行初始化，详情可参考[部署 & 体验](#部署--体验) 。

# ollama 帮助文档
<p align="center" class="flex justify-center">
    <a href="https://www.serverless-devs.com" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=ollama&type=packageType">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=ollama" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=ollama&type=packageVersion">
  </a>
  <a href="http://www.devsapp.cn/details.html?name=ollama" class="ml-1">
    <img src="http://editor.devsapp.cn/icon?package=ollama&type=packageDownload">
  </a>
</p>

<description>

通过函数计算开箱即用Ollama，极低成本部署自己的云端LLM服务。

</description>

<codeUrl>



</codeUrl>
<preview>



</preview>


## 前期准备

使用该项目，您需要有开通以下服务并拥有对应权限：

<service>



| 服务/业务 |  权限  | 相关文档 |
| --- |  --- | --- |
| 函数计算 |  AliyunFCFullAccess | [帮助文档](https://help.aliyun.com/product/2508973.html) [计费文档](https://help.aliyun.com/document_detail/2512928.html) |

</service>

<remark>



</remark>

<disclaimers>



</disclaimers>

## 部署 & 体验

<appcenter>
   
- :fire: 通过 [Serverless 应用中心](https://fcnext.console.aliyun.com/applications/create?template=ollama) ，
  [![Deploy with Severless Devs](https://img.alicdn.com/imgextra/i1/O1CN01w5RFbX1v45s8TIXPz_!!6000000006118-55-tps-95-28.svg)](https://fcnext.console.aliyun.com/applications/create?template=ollama) 该应用。
   
</appcenter>
<deploy>
    
- 通过 [Serverless Devs Cli](https://www.serverless-devs.com/serverless-devs/install) 进行部署：
  - [安装 Serverless Devs Cli 开发者工具](https://www.serverless-devs.com/serverless-devs/install) ，并进行[授权信息配置](https://docs.serverless-devs.com/fc/config) ；
  - 初始化项目：`s init ollama -d ollama`
  - 进入项目，并进行项目部署：`cd ollama && s deploy -y`
   
</deploy>

## 案例介绍

<appdetail id="flushContent">

基于 Ollama 案例，您可以轻松地将 Ollama 服务部署到函数计算中，享受高效、灵活的 AI 语言模型服务。Ollama 是一款专为开发者设计的高性能语言模型，它能够处理各种自然语言处理任务，包括文本生成、翻译、摘要等，是您开发智能应用的强大后盾。

Ollama 以其卓越的性能和易用性在业界获得了广泛认可，Ollama 服务的部署通常需要一定的技术背景和资源配置。然而，通过 Serverless 开发平台，您可以快速、便捷地将 Ollama 服务部署至函数计算，无需担心底层资源管理和运维问题，让您专注于应用的创新和开发。此外，开发平台还提供了包括模型管理、API 集成、云存储等在内的多种服务，以满足不同用户的需求。

使用本案例，servereless开发平台将为您提供开箱即用的ollama服务，并可以**开启GPU闲置预留模式**，在保证使用性能的同时，以最小成本持有自己专属的LLM服务。GPU闲置预留模式公测中，请提交[公测申请](https://survey.aliyun.com/apps/zhiliao/dXfRVPEm-)

当前版本软件内置了以下三种尺寸的Qwen模型：Qwen 0.5b, Qwen 7b, Qwen 14b。每种模型都具有不同的参数量和复杂度，以适应不同的应用场景和性能要求，您可以在应用初始化时进行选择。

</appdetail>

## 使用流程

<usedetail id="flushContent">

### 通过API进行调用
我们推荐以openAI的API范式进行函数调用，其中model需替换为您当前使用的模型名字(`qwen:0.5b`, `qwen:7b`, `qwen:14b`)，如下是一个调用qwen_7b模型的示例：
```
curl "${FunctionEndpoint}/v1/chat/completions" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "qwen:7b",
    "messages": [
      {
        "role": "system",
        "content": "You are a helpful assistant."
      },
      {
        "role": "user",
        "content": "How can you help me?"
      }
    ]
  }'
```
API参数说明及更多请参考 https://github.com/ollama/ollama/blob/main/docs/api.md


### 通过 Open WebUI 进行测试

通过应用 [fc-open-webui](https://github.com/devsapp/fc-open-webui) 部署，并配置 ollama 接口为当前应用触发器

应用中心创建地址 https://fcnext.console.aliyun.com/applications/create?template=fc-open-webui



</usedetail>

## 注意事项

<matters id="flushContent">
</matters>


<devgroup>


## 开发者社区

您如果有关于错误的反馈或者未来的期待，您可以在 [Serverless Devs repo Issues](https://github.com/serverless-devs/serverless-devs/issues) 中进行反馈和交流。如果您想要加入我们的讨论组或者了解 FC 组件的最新动态，您可以通过以下渠道进行：

<p align="center">  

| <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407298906_20211028074819117230.png" width="130px" > | <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407044136_20211028074404326599.png" width="130px" > | <img src="https://serverless-article-picture.oss-cn-hangzhou.aliyuncs.com/1635407252200_20211028074732517533.png" width="130px" > |
| --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| <center>微信公众号：`serverless`</center>                                                                                         | <center>微信小助手：`xiaojiangwh`</center>                                                                                        | <center>钉钉交流群：`33947367`</center>                                                                                           |
</p>
</devgroup>
