/**
 * MCP Server function for Lists entity rows of a particular entity type contained in the request. Note: 1. Currently, only max of one 'sort_by' column is supported. 2. If no 'sort_by' column is provided, the primary key of the table is used. If zero or more than one primary key is available, we default to the unpaginated list entities logic which only returns the first page. 3. The values of the 'sort_by' columns must uniquely identify an entity row, otherwise undefined behaviors may be observed during pagination. 4. Since transactions are not supported, any updates, inserts or deletes during pagination can lead to stale data being returned or other unexpected behaviors.
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

class Get_V2_Parent_EntitiesMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getGet_V2_Parent_EntitiesHandler(MCPServer.APIConfig config) {
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
        if (args.containsKey("parent")) {
            queryParams.add("parent=" + args.get("parent"));
        }
        if (args.containsKey("conditions")) {
            queryParams.add("conditions=" + args.get("conditions"));
        }
        if (args.containsKey("pageToken")) {
            queryParams.add("pageToken=" + args.get("pageToken"));
        }
        if (args.containsKey("pageSize")) {
            queryParams.add("pageSize=" + args.get("pageSize"));
        }
        if (args.containsKey("prettyPrint")) {
            queryParams.add("prettyPrint=" + args.get("prettyPrint"));
        }
        if (args.containsKey("sortBy")) {
            queryParams.add("sortBy=" + args.get("sortBy"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/get_v2_parent_entities" + queryString;
                
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
    
    public static MCPServer.Tool createGet_V2_Parent_EntitiesTool(MCPServer.APIConfig config) {
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
        Map<String, Object> parentProperty = new HashMap<>();
        parentProperty.put("type", "string");
        parentProperty.put("required", true);
        parentProperty.put("description", "Required. Resource name of the Entity Type. Format: projects/{project}/locations/{location}/connections/{connection}/entityTypes/{type}");
        properties.put("parent", parentProperty);
        Map<String, Object> conditionsProperty = new HashMap<>();
        conditionsProperty.put("type", "string");
        conditionsProperty.put("required", false);
        conditionsProperty.put("description", "Conditions to be used when listing entities. From a proto standpoint, There are no restrictions on what can be passed using this field. The connector documentation should have information about what format of filters/conditions are supported.");
        properties.put("conditions", conditionsProperty);
        Map<String, Object> pageTokenProperty = new HashMap<>();
        pageTokenProperty.put("type", "string");
        pageTokenProperty.put("required", false);
        pageTokenProperty.put("description", "Page token value if available from a previous request.");
        properties.put("pageToken", pageTokenProperty);
        Map<String, Object> pageSizeProperty = new HashMap<>();
        pageSizeProperty.put("type", "string");
        pageSizeProperty.put("required", false);
        pageSizeProperty.put("description", "Number of entity rows to return. Defaults page size = 25. Max page size = 200.");
        properties.put("pageSize", pageSizeProperty);
        Map<String, Object> prettyPrintProperty = new HashMap<>();
        prettyPrintProperty.put("type", "string");
        prettyPrintProperty.put("required", false);
        prettyPrintProperty.put("description", "Returns response with indentations and line breaks.");
        properties.put("prettyPrint", prettyPrintProperty);
        Map<String, Object> sortByProperty = new HashMap<>();
        sortByProperty.put("type", "string");
        sortByProperty.put("required", false);
        sortByProperty.put("description", "List of 'sort_by' columns to use when returning the results.");
        properties.put("sortBy", sortByProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "get_v2_parent_entities",
            "Lists entity rows of a particular entity type contained in the request. Note: 1. Currently, only max of one 'sort_by' column is supported. 2. If no 'sort_by' column is provided, the primary key of the table is used. If zero or more than one primary key is available, we default to the unpaginated list entities logic which only returns the first page. 3. The values of the 'sort_by' columns must uniquely identify an entity row, otherwise undefined behaviors may be observed during pagination. 4. Since transactions are not supported, any updates, inserts or deletes during pagination can lead to stale data being returned or other unexpected behaviors.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getGet_V2_Parent_EntitiesHandler(config));
    }
    
}