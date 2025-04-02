import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select"
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table"
import { PlusCircle, Search, Filter, MoreHorizontal, CheckCircle, Clock, AlertCircle } from "lucide-react"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { Badge } from "@/components/ui/badge"

export default function TicketsPage() {
  return (
    <div className="flex-1 p-8 space-y-6">
      <div className="flex items-center justify-between">
        <h1 className="text-3xl font-bold tracking-tight">Tickets</h1>
        <Button>
          <PlusCircle className="mr-2 h-4 w-4" />
          New Ticket
        </Button>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>All Tickets</CardTitle>
          <CardDescription>View and manage all support tickets</CardDescription>
        </CardHeader>
        <CardContent>
          <div className="flex flex-col md:flex-row gap-4 mb-6">
            <div className="relative flex-1">
              <Search className="absolute left-2.5 top-2.5 h-4 w-4 text-muted-foreground" />
              <Input placeholder="Search tickets..." className="pl-8" />
            </div>
            <div className="flex gap-4">
              <Select defaultValue="all">
                <SelectTrigger className="w-[180px]">
                  <SelectValue placeholder="Status" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="all">All Statuses</SelectItem>
                  <SelectItem value="open">Open</SelectItem>
                  <SelectItem value="in-progress">In Progress</SelectItem>
                  <SelectItem value="resolved">Resolved</SelectItem>
                </SelectContent>
              </Select>
              <Select defaultValue="all">
                <SelectTrigger className="w-[180px]">
                  <SelectValue placeholder="Priority" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="all">All Priorities</SelectItem>
                  <SelectItem value="low">Low</SelectItem>
                  <SelectItem value="medium">Medium</SelectItem>
                  <SelectItem value="high">High</SelectItem>
                </SelectContent>
              </Select>
              <Button variant="outline" size="icon">
                <Filter className="h-4 w-4" />
              </Button>
            </div>
          </div>

          <div className="rounded-md border">
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead className="w-[100px]">ID</TableHead>
                  <TableHead>Subject</TableHead>
                  <TableHead>Requester</TableHead>
                  <TableHead>Status</TableHead>
                  <TableHead>Priority</TableHead>
                  <TableHead>Assigned To</TableHead>
                  <TableHead className="text-right">Actions</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                {[
                  {
                    id: "T-1001",
                    subject: "Forgot login password for laptop",
                    requester: "john.doe@example.com",
                    status: "open",
                    priority: "medium",
                    assignedTo: "admin",
                  },
                  {
                    id: "T-1002",
                    subject: "VPN access not working",
                    requester: "sarah.smith@example.com",
                    status: "in-progress",
                    priority: "high",
                    assignedTo: "admin",
                  },
                  {
                    id: "T-1003",
                    subject: "Need software installation",
                    requester: "mike.johnson@example.com",
                    status: "open",
                    priority: "low",
                    assignedTo: "unassigned",
                  },
                  {
                    id: "T-1004",
                    subject: "Email configuration issue",
                    requester: "lisa.wong@example.com",
                    status: "in-progress",
                    priority: "medium",
                    assignedTo: "admin",
                  },
                  {
                    id: "T-1005",
                    subject: "Printer not connecting",
                    requester: "robert.chen@example.com",
                    status: "resolved",
                    priority: "medium",
                    assignedTo: "admin",
                  },
                ].map((ticket) => (
                  <TableRow key={ticket.id}>
                    <TableCell className="font-medium">{ticket.id}</TableCell>
                    <TableCell>{ticket.subject}</TableCell>
                    <TableCell>{ticket.requester}</TableCell>
                    <TableCell>
                      <Badge
                        variant="outline"
                        className={`
                          ${ticket.status === "open" ? "border-blue-500 text-blue-500 bg-blue-50" : ""}
                          ${ticket.status === "in-progress" ? "border-yellow-500 text-yellow-500 bg-yellow-50" : ""}
                          ${ticket.status === "resolved" ? "border-green-500 text-green-500 bg-green-50" : ""}
                        `}
                      >
                        {ticket.status === "open" && <AlertCircle className="mr-1 h-3 w-3" />}
                        {ticket.status === "in-progress" && <Clock className="mr-1 h-3 w-3" />}
                        {ticket.status === "resolved" && <CheckCircle className="mr-1 h-3 w-3" />}
                        {ticket.status.charAt(0).toUpperCase() + ticket.status.slice(1).replace("-", " ")}
                      </Badge>
                    </TableCell>
                    <TableCell>
                      <Badge
                        variant="outline"
                        className={`
                          ${ticket.priority === "low" ? "border-green-500 text-green-500 bg-green-50" : ""}
                          ${ticket.priority === "medium" ? "border-yellow-500 text-yellow-500 bg-yellow-50" : ""}
                          ${ticket.priority === "high" ? "border-red-500 text-red-500 bg-red-50" : ""}
                        `}
                      >
                        {ticket.priority.charAt(0).toUpperCase() + ticket.priority.slice(1)}
                      </Badge>
                    </TableCell>
                    <TableCell>
                      {ticket.assignedTo === "unassigned" ? (
                        <span className="text-gray-500">Unassigned</span>
                      ) : (
                        ticket.assignedTo
                      )}
                    </TableCell>
                    <TableCell className="text-right">
                      <DropdownMenu>
                        <DropdownMenuTrigger asChild>
                          <Button variant="ghost" className="h-8 w-8 p-0">
                            <span className="sr-only">Open menu</span>
                            <MoreHorizontal className="h-4 w-4" />
                          </Button>
                        </DropdownMenuTrigger>
                        <DropdownMenuContent align="end">
                          <DropdownMenuLabel>Actions</DropdownMenuLabel>
                          <DropdownMenuItem>View details</DropdownMenuItem>
                          <DropdownMenuItem>Assign to me</DropdownMenuItem>
                          <DropdownMenuSeparator />
                          <DropdownMenuItem>Change status</DropdownMenuItem>
                          <DropdownMenuItem>Change priority</DropdownMenuItem>
                          <DropdownMenuSeparator />
                          <DropdownMenuItem>Resolve ticket</DropdownMenuItem>
                        </DropdownMenuContent>
                      </DropdownMenu>
                    </TableCell>
                  </TableRow>
                ))}
              </TableBody>
            </Table>
          </div>

          <div className="flex items-center justify-between mt-4">
            <div className="text-sm text-muted-foreground">Showing 5 of 42 tickets</div>
            <div className="flex items-center space-x-2">
              <Button variant="outline" size="sm" disabled>
                Previous
              </Button>
              <Button variant="outline" size="sm">
                Next
              </Button>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}

