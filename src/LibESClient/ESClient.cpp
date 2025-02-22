
#include "ESClient.h"
#include "util/tc_http.h"

void ESClient::setAddresses(const vector<std::tuple<std::string, int>>& addresses, const std::string& protocol)
{
    std::ostringstream os;
    os << "es@";
    for (size_t i = 0; i < addresses.size(); ++i)
    {
        if (i != 0)
        {
            os << ";";
        }
        if ("http" == protocol)
        {
            os << "tcp -h " << std::get<0>(addresses[i]) << " -p " << std::get<1>(addresses[i]);
        }
        else if ("https" == protocol)
        {
            os << "ssl -h " << std::get<0>(addresses[i]) << " -p " << std::get<1>(addresses[i]);
        }
        else
        {
            TLOGERROR("bad es protocol: " << protocol << endl);
            std::cout << "bad es protocol: " << protocol << std::endl;
            throw std::runtime_error("bad es protocol");
        }
    }
    _esPrx = Application::getCommunicator()->stringToProxy<ServantPrx>(os.str());
    _esPrx->tars_set_protocol(ServantProxy::PROTOCOL_HTTP1, 5);
}

int ESClient::doRequest(ESClientRequestMethod method, const std::string& url, const std::string& body, std::string& response)
{
    auto request = std::make_shared<TC_HttpRequest>();
    switch (method)
    {
    case ESClientRequestMethod::Post:
    {
        request->setHeader("Content-Type", "application/x-ndjson");
        request->setPostRequest(url, body);
    }
        break;
    case ESClientRequestMethod::Put:
    {
        request->setHeader("Content-Type", "application/json");
        request->setPutRequest(url, body);
    }
        break;
    default:
    case ESClientRequestMethod::Get:
    {
        //tc_http GET Method Not Support Body ,So We Use POST Instead.
        request->setHeader("Content-Type", "application/json");
        request->setPostRequest(url, body);
    }
        break;
    }
    auto pResponse = std::make_shared<TC_HttpResponse>();
    try
    {
        _esPrx->http_call("es", request, pResponse);
    }
    catch (const std::exception& e)
    {
        response = e.what();
        return -1;
    }
    response = pResponse->getContent();
    return pResponse->getStatus();
}
