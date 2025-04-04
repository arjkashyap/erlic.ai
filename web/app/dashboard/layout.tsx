"use client"

import type React from "react"

import Link from "next/link"
import { Button } from "@/components/ui/button"
import {
  LayoutDashboard,
  TicketIcon,
  MessageSquare,
  Settings,
  Users,
  LogOut,
  Menu,
  Shield,
  Bell,
  Search,
} from "lucide-react"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Input } from "@/components/ui/input"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { useState } from "react"
import { useMobile } from "@/hooks/use-mobile"
import { useCurrentUser } from "@/hooks/use-current-user"
import { useRouter } from "next/navigation"

export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode
}) {
  const isMobile = useMobile()
  const [sidebarOpen, setSidebarOpen] = useState(!isMobile)
  const { user, loading } = useCurrentUser()
  const router = useRouter()

  const handleLogout = async () => {
    try {
      const response = await fetch("/api/auth/logout", {
        method: "POST",
        credentials: "include",
      })

      if (response.ok) {
        router.push("/")
      }
    } catch (error) {
      console.error("Error logging out:", error)
    }
  }

  console.log("user from the layout: " + user)


  return (
    <div className="min-h-screen bg-gray-100">
      {/* Top Navigation */}
      <header className="bg-white border-b sticky top-0 z-30">
        <div className="flex h-16 items-center px-4 md:px-6">
          <div className="flex items-center">
            <Button variant="ghost" size="icon" className="md:hidden mr-2" onClick={() => setSidebarOpen(!sidebarOpen)}>
              <Menu className="h-6 w-6" />
            </Button>
            <Link href="/" className="flex items-center space-x-2">
              <Shield className="h-6 w-6 text-purple-700" />
              <span className="text-xl font-bold hidden md:inline-block">Erlic</span>
            </Link>
          </div>

          <div className="ml-auto flex items-center space-x-4">
            <div className="relative hidden md:block">
              <Search className="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
              <Input type="search" placeholder="Search..." className="w-[200px] lg:w-[300px] pl-8" />
            </div>

            <Button variant="ghost" size="icon">
              <Bell className="h-5 w-5" />
            </Button>

            <DropdownMenu>
              <DropdownMenuTrigger asChild>
                <Button variant="ghost" className="relative h-8 w-8 rounded-full">
                  <Avatar className="h-8 w-8 bg-gray-900">
                    <AvatarFallback className="text-black">
                      {user?.Username ? user.Username[0].toUpperCase() : "U"}
                    </AvatarFallback>
                  </Avatar>
                </Button>
              </DropdownMenuTrigger>
              <DropdownMenuContent className="w-56" align="end" forceMount>
                <DropdownMenuLabel className="font-normal">
                  <div className="flex flex-col space-y-1">
                    <p className="text-sm font-medium leading-none">
                      {loading ? "Loading..." : user?.FullName || "User"}
                    </p>
                    <p className="text-xs leading-none text-muted-foreground">
                      {loading ? "Loading..." : user?.Email || "No email"}
                    </p>
                  </div>
                </DropdownMenuLabel>
                <DropdownMenuSeparator />
                <DropdownMenuItem>
                  <Users className="mr-2 h-4 w-4" />
                  <span>My Team</span>
                </DropdownMenuItem>
                <DropdownMenuItem>
                  <Settings className="mr-2 h-4 w-4" />
                  <span>Settings</span>
                </DropdownMenuItem>
                <DropdownMenuSeparator />
                <DropdownMenuItem onClick={handleLogout}>
                  <LogOut className="mr-2 h-4 w-4" />
                  <span>Log out</span>
                </DropdownMenuItem>
              </DropdownMenuContent>
            </DropdownMenu>
          </div>
        </div>
      </header>

      <div className="flex">
        {/* Sidebar */}
        <aside
          className={`${sidebarOpen ? "block" : "hidden"} md:block bg-white border-r w-64 fixed h-[calc(100vh-4rem)] z-20`}
        >
          <div className="flex flex-col h-full">
            <div className="flex-1 py-6 px-4 space-y-1">
              <Link href="/dashboard">
                <Button variant="ghost" className="w-full justify-start">
                  <LayoutDashboard className="mr-2 h-5 w-5" />
                  Dashboard
                </Button>
              </Link>
              <Link href="/dashboard/tickets">
                <Button variant="ghost" className="w-full justify-start">
                  <TicketIcon className="mr-2 h-5 w-5" />
                  Tickets
                </Button>
              </Link>
              <Link href="/dashboard/ai-assistant">
                <Button variant="ghost" className="w-full justify-start">
                  <MessageSquare className="mr-2 h-5 w-5" />
                  AI Assistant
                </Button>
              </Link>
              <Link href="/dashboard/connectors">
                <Button variant="ghost" className="w-full justify-start">
                  <Shield className="mr-2 h-5 w-5" />
                  Connectors
                </Button>
              </Link>
              <Link href="/dashboard/team">
                <Button variant="ghost" className="w-full justify-start">
                  <Users className="mr-2 h-5 w-5" />
                  Team
                </Button>
              </Link>
            </div>
            <div className="p-4 border-t">
              <Link href="/dashboard/settings">
                <Button variant="ghost" className="w-full justify-start">
                  <Settings className="mr-2 h-5 w-5" />
                  Settings
                </Button>
              </Link>
            </div>
          </div>
        </aside>

        {/* Main Content */}
        <main className={`flex-1 ${sidebarOpen ? "md:ml-64" : ""}`}>{children}</main>
      </div>
    </div>
  )
}

