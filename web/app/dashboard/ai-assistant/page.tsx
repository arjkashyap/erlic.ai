import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { MessageSquare, Send, Settings, AlertCircle, CheckCircle } from "lucide-react"

export default function AIAssistantPage() {
  return (
    <div className="flex-1 p-8 space-y-6">
      <div className="flex items-center justify-between">
        <h1 className="text-3xl font-bold tracking-tight">AI Assistant</h1>
        <Button variant="outline">
          <Settings className="mr-2 h-4 w-4" />
          Configure
        </Button>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        {/* Chat Interface */}
        <div className="md:col-span-2 flex flex-col h-[calc(100vh-12rem)]">
          <Card className="flex-1 flex flex-col">
            <CardHeader className="border-b">
              <CardTitle>IT Assistant</CardTitle>
              <CardDescription>Use natural language to automate IT tasks</CardDescription>
            </CardHeader>
            <CardContent className="flex-1 overflow-auto p-0">
              <div className="p-4 space-y-4">
                {/* System Message */}
                <div className="flex items-start gap-3">
                  <div className="bg-purple-100 p-2 rounded-full">
                    <MessageSquare className="h-5 w-5 text-purple-700" />
                  </div>
                  <div className="bg-gray-100 rounded-lg p-3 max-w-[80%]">
                    <p className="text-sm">
                      Hello! I'm your IT assistant. I can help you automate tasks like resetting passwords, managing
                      accounts, and configuring systems. What would you like to do today?
                    </p>
                  </div>
                </div>

                {/* User Message */}
                <div className="flex items-start gap-3 justify-end">
                  <div className="bg-purple-700 text-white rounded-lg p-3 max-w-[80%]">
                    <p className="text-sm">Reset AD account for username: arjkashyap</p>
                  </div>
                  <div className="bg-gray-200 p-2 rounded-full">
                    <div className="h-5 w-5 rounded-full bg-gray-400" />
                  </div>
                </div>

                {/* System Response */}
                <div className="flex items-start gap-3">
                  <div className="bg-purple-100 p-2 rounded-full">
                    <MessageSquare className="h-5 w-5 text-purple-700" />
                  </div>
                  <div className="bg-gray-100 rounded-lg p-3 max-w-[80%]">
                    <p className="text-sm">
                      I'll help you reset the Active Directory account for user <strong>arjkashyap</strong>. Here are
                      the user details:
                    </p>
                    <div className="mt-2 p-2 bg-white rounded border">
                      <p className="text-xs">
                        <strong>Username:</strong> arjkashyap
                      </p>
                      <p className="text-xs">
                        <strong>Full Name:</strong> Arjun Kashyap
                      </p>
                      <p className="text-xs">
                        <strong>Email:</strong> arjun.kashyap@example.com
                      </p>
                      <p className="text-xs">
                        <strong>Department:</strong> Engineering
                      </p>
                    </div>
                    <p className="text-sm mt-2">Do you want to reset the password for this account?</p>
                    <div className="mt-2 flex space-x-2">
                      <Button size="sm" className="text-xs">
                        Yes, reset password
                      </Button>
                      <Button size="sm" variant="outline" className="text-xs">
                        Cancel
                      </Button>
                    </div>
                  </div>
                </div>

                {/* System Confirmation */}
                <div className="flex items-start gap-3">
                  <div className="bg-purple-100 p-2 rounded-full">
                    <MessageSquare className="h-5 w-5 text-purple-700" />
                  </div>
                  <div className="bg-gray-100 rounded-lg p-3 max-w-[80%]">
                    <div className="flex items-center gap-2 text-green-600 mb-2">
                      <CheckCircle className="h-4 w-4" />
                      <p className="text-sm font-medium">Password reset successful</p>
                    </div>
                    <p className="text-sm">
                      I've reset the password for <strong>arjkashyap</strong>. The temporary password is:{" "}
                      <code>Temp1234!</code>
                    </p>
                    <p className="text-sm mt-2">
                      The user will be prompted to change this password on next login. Is there anything else you'd like
                      me to do?
                    </p>
                  </div>
                </div>
              </div>
            </CardContent>
            <div className="p-4 border-t">
              <div className="flex gap-2">
                <Input placeholder="Type a command or question..." className="flex-1" />
                <Button>
                  <Send className="h-4 w-4" />
                </Button>
              </div>
              <div className="mt-2 text-xs text-muted-foreground">
                <p>Try: "Reset password for user", "Configure Azure connector", "Show open tickets"</p>
              </div>
            </div>
          </Card>
        </div>

        {/* Sidebar */}
        <div className="space-y-6">
          <Card>
            <CardHeader>
              <CardTitle>Connectors</CardTitle>
              <CardDescription>Configure integrations to enable AI actions</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                <div className="flex items-center justify-between p-3 border rounded-lg">
                  <div className="flex items-center gap-3">
                    <div className="bg-green-100 p-2 rounded-full">
                      <CheckCircle className="h-4 w-4 text-green-600" />
                    </div>
                    <div>
                      <p className="font-medium text-sm">Active Directory</p>
                      <p className="text-xs text-muted-foreground">Connected</p>
                    </div>
                  </div>
                  <Button variant="outline" size="sm">
                    Configure
                  </Button>
                </div>

                <div className="flex items-center justify-between p-3 border rounded-lg">
                  <div className="flex items-center gap-3">
                    <div className="bg-yellow-100 p-2 rounded-full">
                      <AlertCircle className="h-4 w-4 text-yellow-600" />
                    </div>
                    <div>
                      <p className="font-medium text-sm">Azure</p>
                      <p className="text-xs text-muted-foreground">Not configured</p>
                    </div>
                  </div>
                  <Button variant="outline" size="sm">
                    Connect
                  </Button>
                </div>

                <div className="flex items-center justify-between p-3 border rounded-lg">
                  <div className="flex items-center gap-3">
                    <div className="bg-yellow-100 p-2 rounded-full">
                      <AlertCircle className="h-4 w-4 text-yellow-600" />
                    </div>
                    <div>
                      <p className="font-medium text-sm">Google Cloud</p>
                      <p className="text-xs text-muted-foreground">Not configured</p>
                    </div>
                  </div>
                  <Button variant="outline" size="sm">
                    Connect
                  </Button>
                </div>
              </div>

              <Button className="w-full mt-4">Add New Connector</Button>
            </CardContent>
          </Card>

          <Card>
            <CardHeader>
              <CardTitle>Command History</CardTitle>
              <CardDescription>Recently executed commands</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-3">
                <div className="p-2 bg-gray-50 rounded-lg text-sm">
                  <p className="font-medium">Reset AD account: arjkashyap</p>
                  <p className="text-xs text-muted-foreground">2 minutes ago</p>
                </div>
                <div className="p-2 bg-gray-50 rounded-lg text-sm">
                  <p className="font-medium">Show open tickets</p>
                  <p className="text-xs text-muted-foreground">Yesterday</p>
                </div>
                <div className="p-2 bg-gray-50 rounded-lg text-sm">
                  <p className="font-medium">Add user to VPN group</p>
                  <p className="text-xs text-muted-foreground">2 days ago</p>
                </div>
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}

