import { rest } from '../rest';
import { MaxTeamPitchingResponse, MinTeamPitchingResponse } from '../type';

const getMaxTeamPitching = async (): Promise<MaxTeamPitchingResponse> => {
  try {
    const { data } = await rest.get<MaxTeamPitchingResponse>('/team/pitching/max');
    return data;
  } catch (error: any) {
    throw new Error(error);
  }
};

const getMinTeamPitching = async (): Promise<MinTeamPitchingResponse> => {
  try {
    const { data } = await rest.get<MinTeamPitchingResponse>('/team/pitching/min');
    return data;
  } catch (error: any) {
    throw new Error(error);
  }
};

export { getMaxTeamPitching, getMinTeamPitching };
