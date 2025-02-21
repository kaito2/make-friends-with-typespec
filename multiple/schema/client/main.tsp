import "@typespec/http";
import "@typespec/openapi3";

using TypeSpec.Http;
using TypeSpec.OpenAPI;

@service({
  title: "Client Service",
})
@server("https://api.example.com")
namespace Client;

model Store {
  name: string;
  address: Address;
}

model Address {
  street: string;
  city: string;
}

@route("/stores")
namespace Stores {
  op list(): Store[];
  op read(@path id: Store): Store;
}