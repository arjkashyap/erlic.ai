import { NextResponse } from "next/server";
import { cookies } from "next/headers";

export async function POST() {
  try {
    const cookieStore = await cookies();
    const sessionCookie = cookieStore.get("_gothic_session");

    if (sessionCookie) {
      // Clear the session cookie
      cookieStore.delete("_gothic_session");
    }

    return NextResponse.json({ success: true });
  } catch (error) {
    console.error("Error in logout:", error);
    return NextResponse.json({ success: false }, { status: 500 });
  }
} 