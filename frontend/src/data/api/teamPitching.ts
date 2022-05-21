import { rest } from '../rest';
import { MaxTeamPitchingResponse, MinTeamPitchingResponse, TeamPitchingResponse } from '../type';

const baseUri = '/team/pitching';

const getMaxTeamPitching = async (): Promise<MaxTeamPitchingResponse> => {
  try {
    const { data } = await rest.get<MaxTeamPitchingResponse>(`${baseUri}/max`);
    return data;
  } catch (error: any) {
    throw new Error(error);
  }
};

const getMinTeamPitching = async (): Promise<MinTeamPitchingResponse> => {
  try {
    const { data } = await rest.get<MinTeamPitchingResponse>(`${baseUri}/min`);
    return data;
  } catch (error: any) {
    throw new Error(error);
  }
};

const getTeamPitching = async (
  teamId: string | undefined,
  year: string
): Promise<TeamPitchingResponse> => {
  try {
    const { data } = await rest.get<TeamPitchingResponse>(`${baseUri}/${teamId}/${year}`);
    return data;
  } catch (error: any) {
    throw new Error(error);
  }
};

const getTeamPitchingByYear = async (fromYear: string, toYear: string): Promise<any> => {
  try {
    return await rest.get<any>(`${baseUri}?from_year=${fromYear}&to_year=${toYear}`);
  } catch (error: any) {
    throw new Error(error);
  }
};

export { getMaxTeamPitching, getMinTeamPitching, getTeamPitching, getTeamPitchingByYear };
