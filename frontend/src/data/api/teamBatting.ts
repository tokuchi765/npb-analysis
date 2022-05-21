import { rest } from '../rest';
import _ from 'lodash';
import { MaxTeamBattingResponse, MinTeamBattingResponse, TeamBattingResponse } from '../type';

const baseUri = '/team/batting';

const getMaxTeamBatting = async (): Promise<MaxTeamBattingResponse> => {
  try {
    const { data } = await rest.get<MaxTeamBattingResponse>(`${baseUri}/max`);
    return data;
  } catch (error: any) {
    throw new Error(error);
  }
};

const getMinTeamBatting = async (): Promise<MinTeamBattingResponse> => {
  try {
    const { data } = await rest.get<MinTeamBattingResponse>(`${baseUri}/min`);
    return data;
  } catch (error: any) {
    throw new Error(error);
  }
};

const getTeamBatting = async (
  teamId: string | undefined,
  year: string
): Promise<TeamBattingResponse> => {
  try {
    const { data } = await rest.get<TeamBattingResponse>(`${baseUri}/${teamId}/${year}`);
    return data;
  } catch (error: any) {
    throw new Error(error);
  }
};

const getTeamBattingByYear = async (fromYear: string, toYear: string): Promise<any> => {
  try {
    return await rest.get<any>(`${baseUri}?from_year=${fromYear}&to_year=${toYear}`);
  } catch (error: any) {
    throw new Error(error);
  }
};

export { getMaxTeamBatting, getMinTeamBatting, getTeamBatting, getTeamBattingByYear };
