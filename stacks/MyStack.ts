import { StackContext, Api } from "@serverless-stack/resources";

export function MyStack({ stack }: StackContext) {
  const api = new Api(stack, "api", {
    routes: {
	  "GET /todos": "handlers/sst/get/main.go",
      "POST /todos": "handlers/sst/create/main.go",
    },
  });
  stack.addOutputs({
    ApiEndpoint: api.url,
  });
}
