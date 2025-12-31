# SAML 服务
SAML（Security Assertion Markup Language，安全断言标记语言）是一种基于 XML 的开放标准，用于在不同安全域之间交换身份验证和授权数据。它主要用于实现单点登录（SSO, Single Sign-On），让用户只需登录一次，即可访问多个相互信任的应用系统。
## SAML 的核心角色
SAML 协议中通常涉及以下三个主要角色：
1. 主体（Subject）
通常是最终用户（例如员工或客户），希望访问某个受保护的资源。

2. 服务提供者（SP, Service Provider）
提供具体服务或应用的系统，如企业邮箱、CRM 系统等。SP 依赖于 IdP 来验证用户身份。

3. 身份提供者（IdP, Identity Provider）
负责对用户进行身份验证，并向 SP 发送包含用户身份信息的 SAML 断言。例如：Okta、Microsoft Entra ID（原 Azure AD）、Keycloak、Shibboleth 等。

## SAML 的工作流程（以 Web SSO 为例）
典型的 SAML SSO 流程如下（SP-initiated）：

1. 用户尝试访问 SP（如公司内部的 HR 系统）。
2. SP 检测到用户未认证，生成一个 SAML 认证请求（AuthnRequest），并重定向用户浏览器到 IdP 的 SSO 服务地址。
3. 用户在 IdP 页面输入凭据（如用户名/密码）进行身份验证。
4. IdP 验证成功后，生成一个包含用户身份信息的 SAML 断言（Assertion），并通过浏览器 POST 回 SP。
5. SP 验证 SAML 断言的签名和有效性，若通过，则创建本地会话，允许用户访问资源。
也存在 IdP-initiated 流程：用户先登录 IdP，然后从 IdP 的仪表板直接跳转到 SP，无需 SP 先发起请求。
## SAML 断言的组成部分
SAML 断言是一个 XML 文档，通常包含以下三种声明之一或多个：

1. 认证声明（Authentication Statement）：说明用户在何时、通过何种方式被认证。
2. 属性声明（Attribute Statement）：包含用户属性（如邮箱、部门、角色等），用于授权决策。
3. 授权决策声明（Authorization Decision Statement）：说明用户是否被允许访问特定资源（较少使用）。