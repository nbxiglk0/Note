- [CI/CD](#cicd)
  - [基础概念](#基础概念)
    - [CI 持续集成（Continuous Integration）](#ci-持续集成continuous-integration)
    - [CD 持续交付（Continuous Delivery）](#cd-持续交付continuous-delivery)
    - [CD 持续部署（Continuous Deployment）](#cd-持续部署continuous-deployment)
  - [常见工具](#常见工具)
  - [CI/CD应用场景](#cicd应用场景)

# CI/CD
## 基础概念
CI/CD 的核心概念是持续集成、持续交付和持续部署。这些关联的事务通常被统称为 CI/CD 管道，由开发和运维团队以敏捷方式协同支持。
### CI 持续集成（Continuous Integration）
持续集成（CI）可以帮助开发者更加方便地将代码更改合并到主分支。  
一旦开发人员将改动的代码合并到主分支，系统就会通过自动构建应用，这个过程通常指使用构建工具（如 Maven、Gradle 等）对代码进行编译、打包和生成可执行文件等操作。构建工具需要与版本控制系统配合使用，当有新的代码提交时，构建工具会自动触发构建操作。  

并运行不同级别的自动化测试（通常是单元测试和集成测试）来验证这些更改，这个过程通常指使用测试框架（如 JUnit、TestNG 等）对代码进行单元测试、集成测试、功能测试和性能测试等操作。确保这些更改没有对应用造成破坏。  
如果自动化测试发现新代码和现有代码之间存在冲突，CI 可以更加轻松地快速修复这些错误。
### CD 持续交付（Continuous Delivery）
CI 在完成了构建、单元测试和集成测试这些自动化流程后，持续交付可以自动把已验证的代码发布到企业自己的存储库。  
持续交付旨在建立一个可随时将开发环境的功能部署到生产环境的代码库。  
在持续交付过程中，每个步骤都涉及到了测试自动化和代码发布自动化。
在流程结束时，运维团队可以快速、轻松地将应用部署到生产环境中。
### CD 持续部署（Continuous Deployment）
它是作为持续交付的延伸，持续部署可以自动将应用发布到生产环境。这个过程通常指使用部署工具（如 Ansible、Puppet、Chef 等）将构建好的软件包部署到目标环境中。  
## 常见工具
**CI 工具**：常见的 CI 工具包括 Jenkins、GitLab CI、Travis CI、CircleCI 等。它们可以与 GitLab、GitHub、Bitbucket 等版本控制系统集成使用。  
**CD 工具**：常见的 CD 工具包括 Ansible、Puppet、Chef 等。它们可以自动化部署和管理基础设施以及应用程序。  
**测试框架**：常见的测试框架包括 JUnit、TestNG、Selenium、JMeter 等。它们可以对代码进行单元测试、集成测试、功能测试和性能测试等操作。  
## CI/CD应用场景
1. 开发人员将本地代码上传gitlab版本服务器
2. jenkins通过webhook插件自动到gitlab服务器拉取最新代码
3. 通过docker-maven-plugin插件自动编译代码
4. 将自定义镜像上传docker私服仓库
5. k8s集群自动拉取最新版本镜像
6. 自动化部署整个项目
7. 用户通过nginx负载均衡访问整个项目