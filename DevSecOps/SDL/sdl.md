- [SDLC](#sdlc)
  - [三方库漏洞检测](#三方库漏洞检测)
    - [获得应用所依赖的三方库](#获得应用所依赖的三方库)
      - [二次依赖问题](#二次依赖问题)
    - [获得漏洞库](#获得漏洞库)
    - [对比得出结果](#对比得出结果)
      - [和CI/CD流程集成](#和cicd流程集成)
  - [检测工具](#检测工具)
    - [OWASP-Dependency-Check](#owasp-dependency-check)
  - [参考](#参考)

# SDLC
## 三方库漏洞检测
实现机制：应用所依赖的三方库与漏洞库（cve、cnvd等）做对比，看是否有使用有漏洞的三方库版本。
分为三个步骤：
*  获得应用所依赖的三方库。
*  获得漏洞库。
*  做对比得出结果并通知修复。 
### 获得应用所依赖的三方库
应用程序如果使用开源三方库，一般会在配置文件中列出依赖包及其版本。所以最简单的方式就是去读取分析这个配置文件来获取依赖项信息。   
不同的语言有不同的配置文件：  
```
golang：go.sum
java：pom.xml，bulid.gradle
python：requirements.txt
nodejs：package.json
php：composer.lock 
```
#### 二次依赖问题
三方库是最新的，但是三方库的依赖有旧的有漏洞的三方库。  
解决方式：
1. 在业务代码之外维护一个公开项目，手动更新或者修复第三方依赖的问题，并定期用脚本对比自己维护的代码与第三方代码的差异，及时在第三方修复bug或者漏洞后把依赖再迁移回去。
### 获得漏洞库
定时拉取CVE,CNVD数据库，并实现一个三方库漏洞检测服务，当传入三方库的信息时能够返回是否有漏洞存在。
### 对比得出结果
检测时间最好在平时commit时就触发检测，如果在build的时候才检测留给研发再次修改的时间会比较紧，当使用了有漏洞的三方库时将消息推送到对应的开发人员，根据实际情况可以决定是否中断CI/CD流程。
#### 和CI/CD流程集成
CICD是基于自动化脚本的，我们需要将相关检测流程根据开发使用的CI/CD平台工具进行脚本化，然后插入当其自动触发流程中。  
以gitlib为例：  
* pipeline：是一个概念—任务流，没有具体的实体。构建中的阶段（stages）集合，比如自动构建、自动进行单元测试、代码审计等等，会按照顺序执行，所有阶段（stages）执行成功后，才算构建任务（pipeline）执行成功；如果某一个stage失败，后续不再执行，构建任务失败；而一个阶段（stage）可以包含多个job，这些job可以并行执行，某个失败即stage失败；这些stages、job都是定义在.gitlab-ci.yml中的。  
 ![](2023-09-04-16-00-34.png)  
* runner：jobs的执行器。参考：https://www.cnblogs.com/cnundefined/p/7095368.html    
* .gitlab-ci.yml：用来指定构建、测试和部署流程、以及CI触发条件的脚本。Gitlab检测到.gitlab-ci.yml文件，若当前提交（commit）符合文件中指定的触发条件，则会使用配置的gitlab-runner服务运行该脚本进行测试等工作。
## 检测工具
### OWASP-Dependency-Check
参考：https://developer.aliyun.com/article/698621
## 参考
https://www.bilibili.com/read/cv10419374/  
https://www.bilibili.com/read/cv11528596/?spm_id_from=333.999.0.0