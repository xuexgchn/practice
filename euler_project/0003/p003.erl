%%%-------------------------------------------------------------------
%%% @author xuexg
%%% @doc
%%%     ProjectEuler Problem 3
%%% @end
%%% Created : 17. 十月 2017 15:01
%%%-------------------------------------------------------------------
-module(p003).

%% API
-export([solve/1]).

is_prime(Num) ->
    is_prime(Num, 2).
is_prime(Num, X) when X < Num ->
    case Num rem X of
        0 -> throw({error, not_prime});
        _ -> is_prime(Num, X + 1)
    end;
is_prime(Num, X) when X >= Num ->
    true.

cal(_Num, _Tmp, Flag) when Flag > 0 ->
    Flag;
cal(Num, Tmp, _Flag) when Tmp > 0 ->
    case Num rem Tmp of
        0 ->
            case catch is_prime(Tmp) of
                true ->
                    cal(Num, Tmp, Tmp);
                {error, _Reason} ->
                    cal(Num, Tmp - 1, -1)
            end;
        _ ->
            cal(Num, Tmp - 1, -1)
    end.

solve(Num) when Num > 0 ->
    Sqrt = round(math:sqrt(Num) + 1),
    Ans = cal(Num, Sqrt, -1),
    io:format("The answer is ~p~n", [Ans]).

%% erlang shell run: p003:solve(13195) = 29
%% erlang shell run: p003:solve(600851475143) = 6857