/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface Any {
  "@type"?: string;
}

export interface Status {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: { "@type"?: string }[];
}

export interface AssetDailyStats {
  amountWithFee?: { denom?: string; amount?: string }[];
  amountNoFee?: { denom?: string; amount?: string }[];
  fee?: { denom?: string; amount?: string }[];

  /** @format int32 */
  countWithFee?: number;

  /** @format int32 */
  countNoFee?: number;
  burned?: { denom?: string; amount?: string }[];

  /** @format uint64 */
  countBurned?: string;
  issued?: { denom?: string; amount?: string }[];

  /** @format uint64 */
  countIssued?: string;
  withdraw?: { denom?: string; amount?: string }[];

  /** @format uint64 */
  countWithdraw?: string;
  refReward?: { denom?: string; amount?: string }[];

  /** @format uint64 */
  countRefReward?: string;
  sysRefReward?: { denom?: string; amount?: string }[];

  /** @format uint64 */
  countSysRefReward?: string;
}

export interface Coin {
  denom?: string;
  amount?: string;
}

export interface FeeDailyStats {
  /** @format uint64 */
  id?: string;
  amountWithFee?: { denom?: string; amount?: string }[];
  amountNoFee?: { denom?: string; amount?: string }[];
  fee?: { denom?: string; amount?: string }[];

  /** @format int32 */
  countWithFee?: number;

  /** @format int32 */
  countNoFee?: number;
}

export interface PageRequest {
  /** @format byte */
  key?: string;

  /** @format uint64 */
  offset?: string;

  /** @format uint64 */
  limit?: string;
  count_total?: boolean;
  reverse?: boolean;
}

export interface PageResponse {
  /** @format byte */
  next_key?: string;

  /** @format uint64 */
  total?: string;
}

export interface QueryAllFeeStatsResponse {
  feeStats?: {
    date?: string;
    index?: string;
    stats?: {
      id?: string;
      amountWithFee?: { denom?: string; amount?: string }[];
      amountNoFee?: { denom?: string; amount?: string }[];
      fee?: { denom?: string; amount?: string }[];
      countWithFee?: number;
      countNoFee?: number;
    };
  }[];
  pagination?: { next_key?: string; total?: string };
}

export interface QueryAssetStatsResponse {
  stats?: {
    dailyStats?: {
      amountWithFee?: { denom?: string; amount?: string }[];
      amountNoFee?: { denom?: string; amount?: string }[];
      fee?: { denom?: string; amount?: string }[];
      countWithFee?: number;
      countNoFee?: number;
      burned?: { denom?: string; amount?: string }[];
      countBurned?: string;
      issued?: { denom?: string; amount?: string }[];
      countIssued?: string;
      withdraw?: { denom?: string; amount?: string }[];
      countWithdraw?: string;
      refReward?: { denom?: string; amount?: string }[];
      countRefReward?: string;
      sysRefReward?: { denom?: string; amount?: string }[];
      countSysRefReward?: string;
    };
  };
}

export interface QueryGetFeeStatsResponse {
  feeStats?: {
    date?: string;
    index?: string;
    stats?: {
      id?: string;
      amountWithFee?: { denom?: string; amount?: string }[];
      amountNoFee?: { denom?: string; amount?: string }[];
      fee?: { denom?: string; amount?: string }[];
      countWithFee?: number;
      countNoFee?: number;
    };
  };
}

export interface QueryParamsResponse {
  params?: object;
}

export interface QueryUserStatsResponse {
  stats?: { countUniqueActiveUsers?: string; countNewUsers?: string };
}

export interface StatsAssetStats {
  dailyStats?: {
    amountWithFee?: { denom?: string; amount?: string }[];
    amountNoFee?: { denom?: string; amount?: string }[];
    fee?: { denom?: string; amount?: string }[];
    countWithFee?: number;
    countNoFee?: number;
    burned?: { denom?: string; amount?: string }[];
    countBurned?: string;
    issued?: { denom?: string; amount?: string }[];
    countIssued?: string;
    withdraw?: { denom?: string; amount?: string }[];
    countWithdraw?: string;
    refReward?: { denom?: string; amount?: string }[];
    countRefReward?: string;
    sysRefReward?: { denom?: string; amount?: string }[];
    countSysRefReward?: string;
  };
}

export interface StatsFeeStats {
  date?: string;
  index?: string;
  stats?: {
    id?: string;
    amountWithFee?: { denom?: string; amount?: string }[];
    amountNoFee?: { denom?: string; amount?: string }[];
    fee?: { denom?: string; amount?: string }[];
    countWithFee?: number;
    countNoFee?: number;
  };
}

export type StatsParams = object;

export interface StatsUserStats {
  /** @format uint64 */
  countUniqueActiveUsers?: string;

  /** @format uint64 */
  countNewUsers?: string;
}

export type MsgUpdateParamsResponse = object;

export type Params = object;

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.instance.defaults.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createFormData(input: Record<string, unknown>): FormData {
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      formData.append(
        key,
        property instanceof Blob
          ? property
          : typeof property === "object" && property !== null
          ? JSON.stringify(property)
          : `${property}`,
      );
      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = (format && this.format) || void 0;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      requestParams.headers.common = { Accept: "*/*" };
      requestParams.headers.post = {};
      requestParams.headers.put = {};

      body = this.createFormData(body as Record<string, unknown>);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
    });
  };
}

/**
 * @title HTTP API Console stwartchain.stats
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryAssetStats
   * @request GET:/stwart/stats/asset/{startDate}/{endDate}
   */
  queryAssetStats = (startDate: string, endDate: string, params: RequestParams = {}) =>
    this.request<
      {
        stats?: {
          dailyStats?: {
            amountWithFee?: { denom?: string; amount?: string }[];
            amountNoFee?: { denom?: string; amount?: string }[];
            fee?: { denom?: string; amount?: string }[];
            countWithFee?: number;
            countNoFee?: number;
            burned?: { denom?: string; amount?: string }[];
            countBurned?: string;
            issued?: { denom?: string; amount?: string }[];
            countIssued?: string;
            withdraw?: { denom?: string; amount?: string }[];
            countWithdraw?: string;
            refReward?: { denom?: string; amount?: string }[];
            countRefReward?: string;
            sysRefReward?: { denom?: string; amount?: string }[];
            countSysRefReward?: string;
          };
        };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/stwart/stats/asset/${startDate}/${endDate}`,
      method: "GET",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryFeeStatsAll
   * @request GET:/stwart/stats/fee_stats
   */
  queryFeeStatsAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<
      {
        feeStats?: {
          date?: string;
          index?: string;
          stats?: {
            id?: string;
            amountWithFee?: { denom?: string; amount?: string }[];
            amountNoFee?: { denom?: string; amount?: string }[];
            fee?: { denom?: string; amount?: string }[];
            countWithFee?: number;
            countNoFee?: number;
          };
        }[];
        pagination?: { next_key?: string; total?: string };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/stwart/stats/fee_stats`,
      method: "GET",
      query: query,
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryFeeStats
   * @request GET:/stwart/stats/fee_stats/{date}
   */
  queryFeeStats = (date: string, params: RequestParams = {}) =>
    this.request<
      {
        feeStats?: {
          date?: string;
          index?: string;
          stats?: {
            id?: string;
            amountWithFee?: { denom?: string; amount?: string }[];
            amountNoFee?: { denom?: string; amount?: string }[];
            fee?: { denom?: string; amount?: string }[];
            countWithFee?: number;
            countNoFee?: number;
          };
        };
      },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/stwart/stats/fee_stats/${date}`,
      method: "GET",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @request GET:/stwart/stats/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<{ params?: object }, { code?: number; message?: string; details?: { "@type"?: string }[] }>({
      path: `/stwart/stats/params`,
      method: "GET",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryUserStats
   * @request GET:/stwart/stats/user/{startDate}/{endDate}
   */
  queryUserStats = (startDate: string, endDate: string, params: RequestParams = {}) =>
    this.request<
      { stats?: { countUniqueActiveUsers?: string; countNewUsers?: string } },
      { code?: number; message?: string; details?: { "@type"?: string }[] }
    >({
      path: `/stwart/stats/user/${startDate}/${endDate}`,
      method: "GET",
      ...params,
    });
}
