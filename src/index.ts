import { Elysia } from "elysia";
import { authController } from "./modules/auth/auth.controller";
import { chatConroller } from "./modules/chat/chat.conroller";

new Elysia()
  .get("/", () => "Hello")
  .use(authController)
  .use(chatConroller)
  .listen(3000);
