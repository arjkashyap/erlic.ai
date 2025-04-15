"use client"

import { ChatInterface } from "@/components/chat-interface"

export default function AIAssistantPage() {
  return (
    <div className="container mx-auto p-6">
      <div className="bg-white rounded-lg shadow-md h-[calc(100vh-12rem)]">
        <ChatInterface />
      </div>
    </div>
  )
}

