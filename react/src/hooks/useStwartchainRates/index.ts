/* eslint-disable @typescript-eslint/no-unused-vars */
import { useQuery, type UseQueryOptions, useInfiniteQuery, type UseInfiniteQueryOptions } from "@tanstack/react-query";
import { useClient } from '../useClient';

export default function useStwartchainRates() {
  const client = useClient();
  const QueryParams = ( options: any) => {
    const key = { type: 'QueryParams',  };    
    return useQuery([key], () => {
      return  client.StwartchainRates.query.queryParams().then( res => res.data );
    }, options);
  }
  
  const QueryAddresses = (id: string,  options: any) => {
    const key = { type: 'QueryAddresses',  id };    
    return useQuery([key], () => {
      const { id } = key
      return  client.StwartchainRates.query.queryAddresses(id).then( res => res.data );
    }, options);
  }
  
  const QueryAddressesAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryAddressesAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.StwartchainRates.query.queryAddressesAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  const QueryRates = (denom: string,  options: any) => {
    const key = { type: 'QueryRates',  denom };    
    return useQuery([key], () => {
      const { denom } = key
      return  client.StwartchainRates.query.queryRates(denom).then( res => res.data );
    }, options);
  }
  
  const QueryRatesAll = (query: any, options: any, perPage: number) => {
    const key = { type: 'QueryRatesAll', query };    
    return useInfiniteQuery([key], ({pageParam = 1}: { pageParam?: number}) => {
      const {query } = key

      query['pagination.limit']=perPage;
      query['pagination.offset']= (pageParam-1)*perPage;
      query['pagination.count_total']= true;
      return  client.StwartchainRates.query.queryRatesAll(query ?? undefined).then( res => ({...res.data,pageParam}) );
    }, {...options,
      getNextPageParam: (lastPage, allPages) => { if ((lastPage.pagination?.total ?? 0) >((lastPage.pageParam ?? 0) * perPage)) {return lastPage.pageParam+1 } else {return undefined}},
      getPreviousPageParam: (firstPage, allPages) => { if (firstPage.pageParam==1) { return undefined } else { return firstPage.pageParam-1}}
    }
    );
  }
  
  return {QueryParams,QueryAddresses,QueryAddressesAll,QueryRates,QueryRatesAll,
  }
}
