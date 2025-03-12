import { useQuery } from "@connectrpc/connect-query";
import { createFileRoute } from "@tanstack/react-router";
import { goat } from "../gen/proto/goat/v1/goat-GoatService_connectquery";

export const Route = createFileRoute("/")({
  component: App,
});

function App() {

  const data = useQuery(goat, { sentence: "Is this the 🐐 stack?" });


  return (
    <div className="text-center">
      <header className="min-h-screen flex flex-col items-center justify-center text-[calc(10px+2vmin)]">

        <p>Is this  the 🐐 stack?</p>

        <p>{data.data?.sentence}</p>
      </header>
    </div>
  );
}
