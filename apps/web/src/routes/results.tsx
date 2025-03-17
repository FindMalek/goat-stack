import { useQuery } from "@connectrpc/connect-query";
import { createFileRoute } from "@tanstack/react-router";
import { getVotes } from "../gen/proto/goat/v1/goat-GoatService_connectquery";
import {
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
  type ChartConfig,
} from "@goat/ui/components/charts";
import { Pie, PieChart } from "recharts";

export const Route = createFileRoute("/results")({
  component: RouteComponent,
});

const chartConfig = {
  yes: {
    label: "Yes",
    color: "var(--chart-1)",
  },
  no: {
    label: "No",
    color: "var(--chart-2)",
  },
} satisfies ChartConfig;

function RouteComponent() {
  const data = useQuery(getVotes, {});

  if (data.isLoading) {
    return null;
  }

  const chartData = [
    {
      name: "Yes",
      value: Number(data.data?.Yes),
      fill: "var(--color-yes)",
    },
    {
      name: "No",
      value: Number(data.data?.No),
      fill: "var(--color-no)",
    },
  ];
  return (
    <div className="flex items-center justify-center flex-col h-screen">
      <h1 className="text-4xl">Is it the goat stack??</h1>

      <ChartContainer config={chartConfig} className="min-h-[100px] min-w-xl">
        <PieChart>
          <Pie dataKey="value" data={chartData} />
          <ChartTooltip content={<ChartTooltipContent />} />
        </PieChart>
      </ChartContainer>
      <p>Star us on <a href="https://github.com/thibaultleouay/goat-stack" className="hover:underline">GitHub</a></p>
    </div>
  );
}
