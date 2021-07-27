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
  // TODO: add this to backend
  online?: boolean;
}

export interface Health {
  id: string;
  agentID: string;
  host: Host;
  createTime: number;
  updateTime: number;
  online?: boolean;
}

export interface HealthResponse extends StandardResponse {
  Data: Health[];
}
