// Generic types

export interface StandardResponse {
  Data: unknown;
  StatusCode: number;
  Error: string;
  Success: boolean;
}

export interface Host {
  hostname?: string;
  uptime?: number;
  bootTime?: number;
  procs?: number;
  os?: string;
  platform?: string;
  platformFamily?: string;
  platformVersion?: string;
  kernelVersion?: string;
  kernelArch?: string;
  virtualizationSystem?: string;
  virtualizationRole?: string;
  hostID?: string;
  online?: boolean;
}

export interface Health {
  id: string;
  agentID: string;
  createTime: number;
  updateTime: number;
  online?: boolean;
}

export interface HealthResponse extends StandardResponse {
  Data: Health[];
}

export interface HostResponse extends StandardResponse {
  Data: Host[];
}
