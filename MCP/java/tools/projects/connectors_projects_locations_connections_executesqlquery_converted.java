/**
 * MCP Server function for Executes a SQL statement specified in the body of the request. An example of this SQL statement in the case of Salesforce connector would be 'select * from Account a, Order o where a.Id = o.AccountId'.
 */

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.*;
import java.util.function.Function;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;

class Post_V2_Connection_Execute_Sql_QueryMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_V2_Connection_Execute_Sql_QueryHandler(MCPServer.APIConfig config) {
        return (request) -> {
            try {
                Map<String, Object> args = request.getArguments();
                if (args == null) {
                    return new MCPServer.MCPToolResult("Invalid arguments object", true);
                }
                
                List<String> queryParams = new ArrayList<>();
        if (args.containsKey("$.xgafv")) {
            queryParams.add("$.xgafv=" + args.get("$.xgafv"));
        }
        if (args.containsKey("access_token")) {
            queryParams.add("access_token=" + args.get("access_token"));
        }
        if (args.containsKey("alt")) {
            queryParams.add("alt=" + args.get("alt"));
        }
        if (args.containsKey("callback")) {
            queryParams.add("callback=" + args.get("callback"));
        }
        if (args.containsKey("fields")) {
            queryParams.add("fields=" + args.get("fields"));
        }
        if (args.containsKey("key")) {
            queryParams.add("key=" + args.get("key"));
        }
        if (args.containsKey("oauth_token")) {
            queryParams.add("oauth_token=" + args.get("oauth_token"));
        }
        if (args.containsKey("quotaUser")) {
            queryParams.add("quotaUser=" + args.get("quotaUser"));
        }
        if (args.containsKey("upload_protocol")) {
            queryParams.add("upload_protocol=" + args.get("upload_protocol"));
        }
        if (args.containsKey("uploadType")) {
            queryParams.add("uploadType=" + args.get("uploadType"));
        }
        if (args.containsKey("connection")) {
            queryParams.add("connection=" + args.get("connection"));
        }
        if (args.containsKey("prettyPrint")) {
            queryParams.add("prettyPrint=" + args.get("prettyPrint"));
        }
        if (args.containsKey("query")) {
            queryParams.add("query=" + args.get("query"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_v2_connection_execute_sql_query" + queryString;
                
                HttpClient client = HttpClient.newHttpClient();
                HttpRequest httpRequest = HttpRequest.newBuilder()
                    .uri(URI.create(url))
                    .header("Authorization", "Bearer " + config.getApiKey())
                    .header("Accept", "application/json")
                    .GET()
                    .build();
                
                HttpResponse<String> response = client.send(httpRequest, HttpResponse.BodyHandlers.ofString());
                
                if (response.statusCode() >= 400) {
                    return new MCPServer.MCPToolResult("API error: " + response.body(), true);
                }
                
                // Pretty print JSON
                ObjectMapper mapper = new ObjectMapper();
                JsonNode jsonNode = mapper.readTree(response.body());
                String prettyJson = mapper.writerWithDefaultPrettyPrinter().writeValueAsString(jsonNode);
                
                return new MCPServer.MCPToolResult(prettyJson);
                
            } catch (IOException | InterruptedException e) {
                return new MCPServer.MCPToolResult("Request failed: " + e.getMessage(), true);
            } catch (Exception e) {
                return new MCPServer.MCPToolResult("Unexpected error: " + e.getMessage(), true);
            }
        };
    }
    
    public static MCPServer.Tool createPost_V2_Connection_Execute_Sql_QueryTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> $.xgafvProperty = new HashMap<>();
        $.xgafvProperty.put("type", "string");
        $.xgafvProperty.put("required", false);
        $.xgafvProperty.put("description", "V1 error format.");
        properties.put("$.xgafv", $.xgafvProperty);
        Map<String, Object> access_tokenProperty = new HashMap<>();
        access_tokenProperty.put("type", "string");
        access_tokenProperty.put("required", false);
        access_tokenProperty.put("description", "OAuth access token.");
        properties.put("access_token", access_tokenProperty);
        Map<String, Object> altProperty = new HashMap<>();
        altProperty.put("type", "string");
        altProperty.put("required", false);
        altProperty.put("description", "Data format for response.");
        properties.put("alt", altProperty);
        Map<String, Object> callbackProperty = new HashMap<>();
        callbackProperty.put("type", "string");
        callbackProperty.put("required", false);
        callbackProperty.put("description", "JSONP");
        properties.put("callback", callbackProperty);
        Map<String, Object> fieldsProperty = new HashMap<>();
        fieldsProperty.put("type", "string");
        fieldsProperty.put("required", false);
        fieldsProperty.put("description", "Selector specifying which fields to include in a partial response.");
        properties.put("fields", fieldsProperty);
        Map<String, Object> keyProperty = new HashMap<>();
        keyProperty.put("type", "string");
        keyProperty.put("required", false);
        keyProperty.put("description", "API key. Your API key identifies your project and provides you with API access, quota, and reports. Required unless you provide an OAuth 2.0 token.");
        properties.put("key", keyProperty);
        Map<String, Object> oauth_tokenProperty = new HashMap<>();
        oauth_tokenProperty.put("type", "string");
        oauth_tokenProperty.put("required", false);
        oauth_tokenProperty.put("description", "OAuth 2.0 token for the current user.");
        properties.put("oauth_token", oauth_tokenProperty);
        Map<String, Object> quotaUserProperty = new HashMap<>();
        quotaUserProperty.put("type", "string");
        quotaUserProperty.put("required", false);
        quotaUserProperty.put("description", "Available to use for quota purposes for server-side applications. Can be any arbitrary string assigned to a user, but should not exceed 40 characters.");
        properties.put("quotaUser", quotaUserProperty);
        Map<String, Object> upload_protocolProperty = new HashMap<>();
        upload_protocolProperty.put("type", "string");
        upload_protocolProperty.put("required", false);
        upload_protocolProperty.put("description", "Upload protocol for media (e.g. \"raw\", \"multipart\\");
        properties.put("upload_protocol", upload_protocolProperty);
        Map<String, Object> uploadTypeProperty = new HashMap<>();
        uploadTypeProperty.put("type", "string");
        uploadTypeProperty.put("required", false);
        uploadTypeProperty.put("description", "Legacy upload protocol for media (e.g. \"media\", \"multipart\\");
        properties.put("uploadType", uploadTypeProperty);
        Map<String, Object> connectionProperty = new HashMap<>();
        connectionProperty.put("type", "string");
        connectionProperty.put("required", true);
        connectionProperty.put("description", "Required. Resource name of the Connection. Format: projects/{project}/locations/{location}/connections/{connection}");
        properties.put("connection", connectionProperty);
        Map<String, Object> prettyPrintProperty = new HashMap<>();
        prettyPrintProperty.put("type", "string");
        prettyPrintProperty.put("required", false);
        prettyPrintProperty.put("description", "Returns response with indentations and line breaks.");
        properties.put("prettyPrint", prettyPrintProperty);
        Map<String, Object> queryProperty = new HashMap<>();
        queryProperty.put("type", "string");
        queryProperty.put("required", false);
        queryProperty.put("description", "Input parameter: A wrapper around the SQL query statement. This is needed so that the JSON representation of ExecuteSqlQueryRequest has the following format: `{\"query\":\"select *\"}`.");
        properties.put("query", queryProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_v2_connection_execute_sql_query",
            "Executes a SQL statement specified in the body of the request. An example of this SQL statement in the case of Salesforce connector would be 'select * from Account a, Order o where a.Id = o.AccountId'.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_V2_Connection_Execute_Sql_QueryHandler(config));
    }
    
}