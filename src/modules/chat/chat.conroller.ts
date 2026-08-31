import { Elysia, t } from "elysia";
import { chatService } from "./chat.service";
import { CreateChatBody, AddParticipantBody } from "./chat.model";

export const chatController = new Elysia({ prefix: "/chat" })
  .get(
    "",
    async ({ headers }) => {
      const userId = headers["x-user-id"];
      return await chatService.getUserChats(userId);
    },
    {
      headers: t.Object({
        "x-user-id": t.String({ description: "ID текущего пользователя" }),
      }),
    },
  )
  .post(
    "/create",
    async ({ body }) => {
      return await chatService.create(body);
    },
    {
      body: CreateChatBody,
    },
  )
  .post(
    "/:chatId/participants",
    async ({ params: { chatId }, body: { userId } }) => {
      return await chatService.addParticipant(chatId, userId);
    },
    {
      params: t.Object({
        chatId: t.String(),
      }),
      body: AddParticipantBody,
    },
  );
