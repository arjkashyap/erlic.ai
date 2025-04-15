import { NextResponse } from "next/server";
import { cookies } from "next/headers";
import { checkAuth } from "@/lib/auth-utils";

export async function POST(request: Request) {
  const cookieStore = await cookies();
  const sessionCookie = cookieStore.get("_gothic_session");
  const authResult = await checkAuth(sessionCookie);

  if (!authResult.authenticated) {
    return NextResponse.json({ error: "Unauthorized" }, { status: 401 });
  }

  try {
    const { message } = await request.json();

    if (!message) {
      return NextResponse.json({ error: "Message is required" }, { status: 400 });
    }

    const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/chat/prompt`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "Cookie": `_gothic_session=${sessionCookie?.value}`,
      },
      body: JSON.stringify({ message }),
    });

    if (!response.ok) {
      throw new Error("Failed to get response from chat service");
    }

    const data = await response.json();
    return NextResponse.json(data);
  } catch (error) {
    console.error("Error in chat prompt:", error);
    return NextResponse.json(
      { error: "Failed to process chat request" },
      { status: 500 }
    );
  }
} 