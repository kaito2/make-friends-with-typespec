import "@typespec/http";
import "@typespec/openapi3";

using TypeSpec.Http;
using TypeSpec.OpenAPI;

@service({
  title: "Admin Service",
})
@server("https://api.example.com")
namespace Admin;

model Client {
  name: string;
}

@route("/clients")
namespace Clients {
  op list(): Client[];
  op read(@path id: Client): Client;
}
