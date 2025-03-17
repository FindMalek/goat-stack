import { useQuery } from '@connectrpc/connect-query';
import { createFileRoute } from '@tanstack/react-router'
import { getVotes } from '../gen/proto/goat/v1/goat-GoatService_connectquery';

export const Route = createFileRoute('/results')({
  component: RouteComponent,
})

function RouteComponent() {

  const data = useQuery(
    getVotes,
    {},
  );

  if(data.isLoading) {
    return <div>Loading...</div>
  }

  return <><div>Yes {data.data?.Yes}</div> <div>No {data.data?.No} </div></>
}
