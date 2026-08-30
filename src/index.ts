import { Elysia } from "elysia";
import { authController } from "./modules/auth/auth.controller";

new Elysia()
  .get("/", () => "Hello")
  .use(authController)
  .listen(3000);
