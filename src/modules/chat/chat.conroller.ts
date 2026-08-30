import { Elysia } from "elysia";
import { chatService } from "./chat.service";
import { CreateChatBody } from "./chat.model";

export const chatConroller = new Elysia({ prefix: "/chat" }).post(
  "/create",
  async ({ body }) => {
    return await chatService.create(body);
  },
  {
    body: CreateChatBody,
  },
);
