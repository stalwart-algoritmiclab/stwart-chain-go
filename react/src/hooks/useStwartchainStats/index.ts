/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/react-query";
import { useClient } from '../useClient';

export default function useStwartchainStats() {
  const client = useClient();
  const QueryAssetStats = (startDate: string, endDate: string,  options: any) => {
    const key = { type: 'QueryAssetStats',  startDate,  endDate };    
    return useQuery([key], () => {
      const { startDate,  endDate } = key
      return  client.StwartchainStats.query.queryAssetStats(startDate, endDate).then( res => res.data );
    }, options);
  }
  
  const QueryUserStats = (startDate: string, endDate: string,  options: any) => {
    const key = { type: 'QueryUserStats',  startDate,  endDate };    
    return useQuery([key], () => {
      const { startDate,  endDate } = key
      return  client.StwartchainStats.query.queryUserStats(startDate, endDate).then( res => res.data );
    }, options);
  }
  
  const QueryParams = ( options: any) => {
    const key = { type: 'QueryParams',  };    
    return useQuery([key], () => {
      return  client.StwartchainStats.query.queryParams().then( res => res.data );
    }, options);
  }
  
  const QueryFeeStats = (date: string,  options: any) => {
    const key = { type: 'QueryFeeStats',  date };    
    return useQuery([key], () => {
      const { date } = key
      return  client.StwartchainStats.query.queryFeeStats(date).then( res => res.data );
    }, options);
  }
  
  const QueryFeeStatsAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryFeeStatsAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.StwartchainStats.query.queryFeeStatsAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  return {QueryAssetStats,QueryUserStats,QueryParams,QueryFeeStats,QueryFeeStatsAll,
  }
}
