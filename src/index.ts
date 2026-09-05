import { Elysia } from "elysia";
import { authController } from "./modules/auth/auth.controller";
import { chatController } from "./modules/chat/chat.controller";
import swagger from "@elysiajs/swagger";
import { userController } from "./modules/user/user.controller";

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
  .use(userController)
  .listen(3000);
