import { CreateChatBody } from "./chat.model";
import { db } from "../../database";
import { table } from "../../database/schema";

export const chatService = {
  async create(data: typeof CreateChatBody.static) {
    const [newChat] = await db
      .insert(table.chat)
      .values({ creatorId: data.creatorId })
      .returning();

    return {
      success: true,
      message: "Чат успешно создан",
      chat: { creatorId: data.creatorId },
    };
  },
};
