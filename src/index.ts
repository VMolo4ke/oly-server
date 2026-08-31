import { Elysia } from "elysia";
import { authController } from "./modules/auth/auth.controller";
import { chatController } from "./modules/chat/chat.conroller";
import swagger from "@elysiajs/swagger";

new Elysia()
  .get("/", () => "Hello")
  .use(
    swagger({
      documentation: {
        info: {
          title: "Oly Server API",
          version: "1.0.0",
          description: "Документация для мессенджера Oly",
        },
      },
      path: "/docs",
    }),
  )
  .use(authController)
  .use(chatController)
  .listen(3000);
