import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { BarChart, TicketIcon, Clock, CheckCircle, AlertCircle, MessageSquare, PlusCircle } from "lucide-react"

export default function DashboardPage() {
  return (
    <div className="flex-1 space-y-6 p-8">
      <div className="flex items-center justify-between">
        <h1 className="text-3xl font-bold tracking-tight">Dashboard</h1>
        <Button>
          <PlusCircle className="mr-2 h-4 w-4" />
          New Ticket
        </Button>
      </div>

      {/* Overview Cards */}
      <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-4">
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Total Tickets</CardTitle>
            <TicketIcon className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">128</div>
            <p className="text-xs text-muted-foreground">+14% from last month</p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Open Tickets</CardTitle>
            <AlertCircle className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">42</div>
            <p className="text-xs text-muted-foreground">-3% from last month</p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Resolved Today</CardTitle>
            <CheckCircle className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">16</div>
            <p className="text-xs text-muted-foreground">+6% from yesterday</p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
            <CardTitle className="text-sm font-medium">Avg. Response Time</CardTitle>
            <Clock className="h-4 w-4 text-muted-foreground" />
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">2.4h</div>
            <p className="text-xs text-muted-foreground">-12% from last month</p>
          </CardContent>
        </Card>
      </div>

      {/* Tabs Section */}
      <Tabs defaultValue="recent" className="space-y-4">
        <TabsList>
          <TabsTrigger value="recent">Recent Tickets</TabsTrigger>
          <TabsTrigger value="assigned">Assigned to Me</TabsTrigger>
          <TabsTrigger value="analytics">Analytics</TabsTrigger>
        </TabsList>

        <TabsContent value="recent" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Recent Tickets</CardTitle>
              <CardDescription>View and manage the most recent support tickets.</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                {/* Ticket Item */}
                {[1, 2, 3, 4, 5].map((ticket) => (
                  <div key={ticket} className="flex items-center justify-between p-4 border rounded-lg">
                    <div className="flex items-center space-x-4">
                      <div
                        className={`w-2 h-10 rounded-full ${ticket % 2 === 0 ? "bg-yellow-500" : "bg-green-500"}`}
                      ></div>
                      <div>
                        <h3 className="font-medium">Forgot login password for laptop</h3>
                        <p className="text-sm text-muted-foreground">Submitted by: user{ticket}@example.com</p>
                      </div>
                    </div>
                    <div className="flex items-center space-x-2">
                      <span className="text-sm text-muted-foreground">2h ago</span>
                      <Button variant="outline" size="sm">
                        View
                      </Button>
                    </div>
                  </div>
                ))}
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="assigned" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Tickets Assigned to Me</CardTitle>
              <CardDescription>Manage tickets that have been assigned to you.</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-4">
                {/* Assigned Ticket Item */}
                {[1, 2, 3].map((ticket) => (
                  <div key={ticket} className="flex items-center justify-between p-4 border rounded-lg">
                    <div className="flex items-center space-x-4">
                      <div className={`w-2 h-10 rounded-full ${ticket === 2 ? "bg-red-500" : "bg-yellow-500"}`}></div>
                      <div>
                        <h3 className="font-medium">
                          {ticket === 1
                            ? "Reset AD account for user"
                            : ticket === 2
                              ? "VPN access not working"
                              : "Email configuration issue"}
                        </h3>
                        <p className="text-sm text-muted-foreground">Priority: {ticket === 2 ? "High" : "Medium"}</p>
                      </div>
                    </div>
                    <div className="flex items-center space-x-2">
                      <Button variant="outline" size="sm">
                        View
                      </Button>
                      <Button size="sm">Resolve</Button>
                    </div>
                  </div>
                ))}
              </div>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="analytics" className="space-y-4">
          <Card>
            <CardHeader>
              <CardTitle>Ticket Analytics</CardTitle>
              <CardDescription>View performance metrics and ticket statistics.</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="h-[300px] flex items-center justify-center border rounded-lg bg-gray-50">
                <div className="text-center">
                  <BarChart className="h-10 w-10 text-muted-foreground mx-auto mb-2" />
                  <p className="text-muted-foreground">Analytics visualization would appear here</p>
                </div>
              </div>
            </CardContent>
          </Card>
        </TabsContent>
      </Tabs>

      {/* AI Assistant Preview */}
      <Card className="border-purple-200 bg-purple-50">
        <CardHeader>
          <CardTitle className="flex items-center">
            <MessageSquare className="h-5 w-5 text-purple-700 mr-2" />
            AI Assistant
          </CardTitle>
          <CardDescription>Use natural language to automate IT tasks</CardDescription>
        </CardHeader>
        <CardContent>
          <div className="bg-white p-4 rounded-lg border border-purple-200">
            <p className="text-sm text-gray-600 mb-4">Try commands like:</p>
            <ul className="space-y-2 text-sm">
              <li className="flex items-center">
                <span className="bg-purple-100 text-purple-800 px-2 py-1 rounded text-xs mr-2">Command</span>
                <span>Reset AD account for username: johndoe</span>
              </li>
              <li className="flex items-center">
                <span className="bg-purple-100 text-purple-800 px-2 py-1 rounded text-xs mr-2">Command</span>
                <span>Configure Windows Active Directory connector</span>
              </li>
              <li className="flex items-center">
                <span className="bg-purple-100 text-purple-800 px-2 py-1 rounded text-xs mr-2">Command</span>
                <span>Show all open tickets assigned to me</span>
              </li>
            </ul>
          </div>
          <div className="mt-4">
            <Button className="w-full">
              <MessageSquare className="mr-2 h-4 w-4" />
              Open AI Assistant
            </Button>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}

