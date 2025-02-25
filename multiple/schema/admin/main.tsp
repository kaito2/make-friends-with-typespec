import "@typespec/http";
import "@typespec/openapi3";

using TypeSpec.Http;
using TypeSpec.OpenAPI;

@service({
  title: "Admin Service",
})
@server("https://api.example.com")
namespace Admin;

model TenantClient {
  name: string;
}

@route("/tenant-clients")
namespace TenantClients {
  op list(): TenantClient[];
  op read(@path id: TenantClient): TenantClient;
}
