import { eq } from "drizzle-orm";
import { db } from "../../database";
import { table } from "../../database/schema";
import { CreateChatBody } from "./chat.model";

export const chatService = {
  async create(data: typeof CreateChatBody.static) {
    return await db.transaction(async (tx) => {
      const [newChat] = await tx
        .insert(table.chat)
        .values({ creatorId: data.creatorId })
        .returning();

      await tx.insert(table.chatParticipant).values({
        chatId: newChat.id,
        userId: data.creatorId,
      });

      return {
        success: true,
        message: "Чат успешно создан",
        chat: newChat,
      };
    });
  },

  async getUserChats(userId: string) {
    const userChats = await db
      .select({
        id: table.chat.id,
        creatorId: table.chat.creatorId,
        createdAt: table.chat.createdAt,
      })
      .from(table.chatParticipant)
      .innerJoin(table.chat, eq(table.chatParticipant.chatId, table.chat.id))
      .where(eq(table.chatParticipant.userId, userId));

    return {
      success: true,
      chats: userChats,
    };
  },

  async addParticipant(chatId: string, userId: string) {
    const [existing] = await db
      .select()
      .from(table.chatParticipant)
      .where(
        and(
          eq(table.chatParticipant.chatId, chatId),
          eq(table.chatParticipant.userId, userId),
        ),
      );

    if (existing) {
      return {
        success: false,
        message: "Пользователь уже является участником этого чата",
      };
    }

    await db.insert(table.chatParticipant).values({
      chatId,
      userId,
    });

    return {
      success: true,
      message: "Пользователь успешно добавлен в чат",
    };
  },
};
