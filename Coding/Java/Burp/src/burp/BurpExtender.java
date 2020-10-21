package burp;

import java.net.URL;
import java.util.ArrayList;
import java.util.List;

public class BurpExtender  implements  IBurpExtender,IScannerCheck{
    private IBurpExtenderCallbacks callbacks;
    private IExtensionHelpers helpers;
    private static final byte[] INJ_TEST = "|".getBytes();
    public void registerExtenderCallbacks (final IBurpExtenderCallbacks callbacks)
    {
// your extension code here
        this.callbacks = callbacks;
        callbacks.setExtensionName("Screen");//插件名字
        IExtensionHelpers helpers = callbacks.getHelpers();
        callbacks.registerExtensionStateListener((IExtensionStateListener) this);
        callbacks.registerScannerCheck(this);
        //helpers.
    }

    @Override
    public List<IScanIssue> doPassiveScan(IHttpRequestResponse baseRequestResponse) {
        return null;
    }

    @Override
    public List<IScanIssue> doActiveScan(IHttpRequestResponse baseRequestResponse, IScannerInsertionPoint insertionPoint) {
        byte[] checkRequest = insertionPoint.buildRequest(INJ_TEST);
        IHttpRequestResponse checkRequestResponse = callbacks.makeHttpRequest(baseRequestResponse.getHttpService(), checkRequest);//生成新的请求
        boolean flag = CheckRemember(checkRequestResponse.getResponse());
        if (flag == true){
            List<IScanIssue> issues = new ArrayList<>(1);
            List<int[]> requestHighlights = new ArrayList<>(1);
            requestHighlights.add(insertionPoint.getPayloadOffsets(INJ_TEST));

            issues.add(new CustomScanIssue(
                    baseRequestResponse.getHttpService(),
                    helpers.analyzeRequest(baseRequestResponse).getUrl(),
                    new IHttpRequestResponse[] { callbacks.applyMarkers(checkRequestResponse,null,null)},
                    "shiro RememberMe",
                    "shiro framework",
                    "information"
            ));

        }

        return null;
    }
    private boolean CheckRemember(byte[] response){
        IResponseInfo responseInfo = helpers.analyzeResponse(response);
        List cookies = responseInfo.getCookies();
        for (int i = 0; i<cookies.size(); i++)
        {
            if (cookies.get(i) == "deleteMe") {
                return true;
            }
        }
        return false;
    }
    @Override
    public int consolidateDuplicateIssues(IScanIssue existingIssue, IScanIssue newIssue) {
        return 0;
    }
}



//
class HttpListener implements IHttpListener{

    @Override
    public void processHttpMessage(int toolFlag, boolean messageIsRequest, IHttpRequestResponse messageInfo) {

    }
}
class Parameter implements IParameter{

    /**
     * This method is used to retrieve the parameter type.
     *
     * @return The parameter type. The available types are defined within this
     * interface.
     */
    @Override
    public byte getType() {
        return 0;
    }

    /**
     * This method is used to retrieve the parameter name.
     *
     * @return The parameter name.
     */
    @Override
    public String getName() {
        return null;
    }

    /**
     * This method is used to retrieve the parameter value.
     *
     * @return The parameter value.
     */
    @Override
    public String getValue() {
        return null;
    }

    /**
     * This method is used to retrieve the start offset of the parameter name
     * within the HTTP request.
     *
     * @return The start offset of the parameter name within the HTTP request,
     * or -1 if the parameter is not associated with a specific request.
     */
    @Override
    public int getNameStart() {
        return 0;
    }

    /**
     * This method is used to retrieve the end offset of the parameter name
     * within the HTTP request.
     *
     * @return The end offset of the parameter name within the HTTP request, or
     * -1 if the parameter is not associated with a specific request.
     */
    @Override
    public int getNameEnd() {
        return 0;
    }

    /**
     * This method is used to retrieve the start offset of the parameter value
     * within the HTTP request.
     *
     * @return The start offset of the parameter value within the HTTP request,
     * or -1 if the parameter is not associated with a specific request.
     */
    @Override
    public int getValueStart() {
        return 0;
    }

    /**
     * This method is used to retrieve the end offset of the parameter value
     * within the HTTP request.
     *
     * @return The end offset of the parameter value within the HTTP request, or
     * -1 if the parameter is not associated with a specific request.
     */
    @Override
    public int getValueEnd() {
        return 0;
    }
}



//漏洞详情类
class CustomScanIssue implements IScanIssue {
    private IHttpService httpService;
    private URL url;
    private IHttpRequestResponse[] httpMessages;
    private String name;
    private String detail;
    private String severity;

    public CustomScanIssue(
            IHttpService httpService,
            URL url,
            IHttpRequestResponse[] httpMessages,
            String name,
            String detail,
            String severity) {
        this.httpService = httpService;
        this.url = url;
        this.httpMessages = httpMessages;
        this.name = name;
        this.detail = detail;
        this.severity = severity;
    }
    @Override
    public URL getUrl()
    {
        return url;
    }

    @Override
    public String getIssueName()
    {
        return name;
    }

    @Override
    public int getIssueType()
    {
        return 0;
    }

    @Override
    public String getSeverity()
    {
        return severity;
    }

    @Override
    public String getConfidence()
    {
        return "Certain";
    }

    @Override
    public String getIssueBackground()
    {
        return null;
    }

    @Override
    public String getRemediationBackground()
    {
        return null;
    }

    @Override
    public String getIssueDetail()
    {
        return detail;
    }

    @Override
    public String getRemediationDetail()
    {
        return null;
    }
    @Override
    public IHttpRequestResponse[] getHttpMessages()
    {
        return httpMessages;
    }

    @Override
    public IHttpService getHttpService()
    {
        return httpService;
    }
}
