/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/react-query";
import { useClient } from '../useClient';

export default function useStwartchainUsers() {
  const client = useClient();
  const QueryParams = ( options: any) => {
    const key = { type: 'QueryParams',  };    
    return useQuery([key], () => {
      return  client.StwartchainUsers.query.queryParams().then( res => res.data );
    }, options);
  }
  
  const QueryStats = (date: string,  options: any) => {
    const key = { type: 'QueryStats',  date };    
    return useQuery([key], () => {
      const { date } = key
      return  client.StwartchainUsers.query.queryStats(date).then( res => res.data );
    }, options);
  }
  
  const QueryStatsAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryStatsAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.StwartchainUsers.query.queryStatsAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryStatsByDate = (startDate: string, endDate: string, query: any, options: any, perPage: number) => {
    const key = { type: 'QueryStatsByDate',  startDate,  endDate, query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const { startDate,  endDate,query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.StwartchainUsers.query.queryStatsByDate(startDate, endDate, query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryUniqueUsers = (date: string,  options: any) => {
    const key = { type: 'QueryUniqueUsers',  date };    
    return useQuery([key], () => {
      const { date } = key
      return  client.StwartchainUsers.query.queryUniqueUsers(date).then( res => res.data );
    }, options);
  }
  
  const QueryUniqueUsersAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryUniqueUsersAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.StwartchainUsers.query.queryUniqueUsersAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryTotal = ( options: any) => {
    const key = { type: 'QueryTotal',  };    
    return useQuery([key], () => {
      return  client.StwartchainUsers.query.queryTotal().then( res => res.data );
    }, options);
  }
  
  return {QueryParams,QueryStats,QueryStatsAll,QueryStatsByDate,QueryUniqueUsers,QueryUniqueUsersAll,QueryTotal,
  }
}
