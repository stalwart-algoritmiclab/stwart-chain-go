/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/react-query";
import { useClient } from '../useClient';

export default function useStwartchainFaucet() {
  const client = useClient();
  const QueryParams = ( options: any) => {
    const key = { type: 'QueryParams',  };    
    return useQuery([key], () => {
      return  client.StwartchainFaucet.query.queryParams().then( res => res.data );
    }, options);
  }
  
  const QueryTokens = (id: string,  options: any) => {
    const key = { type: 'QueryTokens',  id };    
    return useQuery([key], () => {
      const { id } = key
      return  client.StwartchainFaucet.query.queryTokens(id).then( res => res.data );
    }, options);
  }
  
  const QueryTokensAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryTokensAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.StwartchainFaucet.query.queryTokensAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  return {QueryParams,QueryTokens,QueryTokensAll,
  }
}
