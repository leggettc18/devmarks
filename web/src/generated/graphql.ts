import gql from "graphql-tag";
import * as VueApolloComposable from "@vue/apollo-composable";
import * as VueCompositionApi from "vue";
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = {
  [K in keyof T]: T[K];
};
export type MakeOptional<T, K extends keyof T> = Omit<T, K> &
  { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> &
  { [SubKey in K]: Maybe<T[SubKey]> };
export type ReactiveFunction<TParam> = () => TParam;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  Time: any;
};

export type AuthPayload = {
  __typename?: "AuthPayload";
  token?: Maybe<Scalars["String"]>;
  user?: Maybe<User>;
};

export type Bookmark = {
  __typename?: "Bookmark";
  id: Scalars["ID"];
  createdAt: Scalars["Time"];
  updatedAt: Scalars["Time"];
  name: Scalars["String"];
  url: Scalars["String"];
  color?: Maybe<Scalars["String"]>;
  owner: User;
};

export type Mutation = {
  __typename?: "Mutation";
  newBookmark: Bookmark;
  deleteBookmark: Scalars["Boolean"];
  login?: Maybe<AuthPayload>;
  register?: Maybe<AuthPayload>;
  updateBookmark: Bookmark;
};

export type MutationNewBookmarkArgs = {
  name: Scalars["String"];
  url: Scalars["String"];
  color?: Maybe<Scalars["String"]>;
};

export type MutationDeleteBookmarkArgs = {
  id: Scalars["ID"];
};

export type MutationLoginArgs = {
  email: Scalars["String"];
  password: Scalars["String"];
};

export type MutationRegisterArgs = {
  email: Scalars["String"];
  password: Scalars["String"];
};

export type MutationUpdateBookmarkArgs = {
  id: Scalars["ID"];
  name?: Maybe<Scalars["String"]>;
  url?: Maybe<Scalars["String"]>;
  color?: Maybe<Scalars["String"]>;
};

export type Query = {
  __typename?: "Query";
  bookmarks?: Maybe<Array<Bookmark>>;
  me?: Maybe<User>;
};

export type User = {
  __typename?: "User";
  id: Scalars["ID"];
  createdAt: Scalars["Time"];
  updatedAt: Scalars["Time"];
  email: Scalars["String"];
};

export type DeleteBookmarkMutationVariables = Exact<{
  id: Scalars["ID"];
}>;

export type DeleteBookmarkMutation = { __typename?: "Mutation" } & Pick<
  Mutation,
  "deleteBookmark"
>;

export type LoginMutationVariables = Exact<{
  email: Scalars["String"];
  password: Scalars["String"];
}>;

export type LoginMutation = { __typename?: "Mutation" } & {
  login?: Maybe<
    { __typename?: "AuthPayload" } & Pick<AuthPayload, "token"> & {
        user?: Maybe<{ __typename?: "User" } & Pick<User, "id" | "email">>;
      }
  >;
};

export type NewBookmarkMutationVariables = Exact<{
  name: Scalars["String"];
  url: Scalars["String"];
  color?: Maybe<Scalars["String"]>;
}>;

export type NewBookmarkMutation = { __typename?: "Mutation" } & {
  newBookmark: { __typename?: "Bookmark" } & Pick<
    Bookmark,
    "id" | "name" | "url" | "color"
  >;
};

export type RegisterMutationVariables = Exact<{
  email: Scalars["String"];
  password: Scalars["String"];
}>;

export type RegisterMutation = { __typename?: "Mutation" } & {
  register?: Maybe<
    { __typename?: "AuthPayload" } & Pick<AuthPayload, "token"> & {
        user?: Maybe<{ __typename?: "User" } & Pick<User, "id" | "email">>;
      }
  >;
};

export type UpdateBookmarkMutationVariables = Exact<{
  id: Scalars["ID"];
  name?: Maybe<Scalars["String"]>;
  url?: Maybe<Scalars["String"]>;
  color?: Maybe<Scalars["String"]>;
}>;

export type UpdateBookmarkMutation = { __typename?: "Mutation" } & {
  updateBookmark: { __typename?: "Bookmark" } & Pick<
    Bookmark,
    "id" | "name" | "url" | "color"
  >;
};

export type GetBookmarksQueryVariables = Exact<{ [key: string]: never }>;

export type GetBookmarksQuery = { __typename?: "Query" } & {
  bookmarks?: Maybe<
    Array<
      { __typename?: "Bookmark" } & Pick<
        Bookmark,
        "id" | "name" | "url" | "color"
      >
    >
  >;
};

export const DeleteBookmarkDocument = gql`
  mutation deleteBookmark($id: ID!) {
    deleteBookmark(id: $id)
  }
`;

/**
 * __useDeleteBookmarkMutation__
 *
 * To run a mutation, you first call `useDeleteBookmarkMutation` within a Vue component and pass it any options that fit your needs.
 * When your component renders, `useDeleteBookmarkMutation` returns an object that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - Several other properties: https://v4.apollo.vuejs.org/api/use-mutation.html#return
 *
 * @param options that will be passed into the mutation, supported options are listed on: https://v4.apollo.vuejs.org/guide-composable/mutation.html#options;
 *
 * @example
 * const { mutate, loading, error, onDone } = useDeleteBookmarkMutation({
 *   variables: {
 *     id: // value for 'id'
 *   },
 * });
 */
export function useDeleteBookmarkMutation(
  options:
    | VueApolloComposable.UseMutationOptions<
        DeleteBookmarkMutation,
        DeleteBookmarkMutationVariables
      >
    | ReactiveFunction<
        VueApolloComposable.UseMutationOptions<
          DeleteBookmarkMutation,
          DeleteBookmarkMutationVariables
        >
      >
) {
  return VueApolloComposable.useMutation<
    DeleteBookmarkMutation,
    DeleteBookmarkMutationVariables
  >(DeleteBookmarkDocument, options);
}
export type DeleteBookmarkMutationCompositionFunctionResult = VueApolloComposable.UseMutationReturn<
  DeleteBookmarkMutation,
  DeleteBookmarkMutationVariables
>;
export const LoginDocument = gql`
  mutation login($email: String!, $password: String!) {
    login(email: $email, password: $password) {
      token
      user {
        id
        email
      }
    }
  }
`;

/**
 * __useLoginMutation__
 *
 * To run a mutation, you first call `useLoginMutation` within a Vue component and pass it any options that fit your needs.
 * When your component renders, `useLoginMutation` returns an object that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - Several other properties: https://v4.apollo.vuejs.org/api/use-mutation.html#return
 *
 * @param options that will be passed into the mutation, supported options are listed on: https://v4.apollo.vuejs.org/guide-composable/mutation.html#options;
 *
 * @example
 * const { mutate, loading, error, onDone } = useLoginMutation({
 *   variables: {
 *     email: // value for 'email'
 *     password: // value for 'password'
 *   },
 * });
 */
export function useLoginMutation(
  options:
    | VueApolloComposable.UseMutationOptions<
        LoginMutation,
        LoginMutationVariables
      >
    | ReactiveFunction<
        VueApolloComposable.UseMutationOptions<
          LoginMutation,
          LoginMutationVariables
        >
      >
) {
  return VueApolloComposable.useMutation<LoginMutation, LoginMutationVariables>(
    LoginDocument,
    options
  );
}
export type LoginMutationCompositionFunctionResult = VueApolloComposable.UseMutationReturn<
  LoginMutation,
  LoginMutationVariables
>;
export const NewBookmarkDocument = gql`
  mutation newBookmark($name: String!, $url: String!, $color: String) {
    newBookmark(name: $name, url: $url, color: $color) {
      id
      name
      url
      color
    }
  }
`;

/**
 * __useNewBookmarkMutation__
 *
 * To run a mutation, you first call `useNewBookmarkMutation` within a Vue component and pass it any options that fit your needs.
 * When your component renders, `useNewBookmarkMutation` returns an object that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - Several other properties: https://v4.apollo.vuejs.org/api/use-mutation.html#return
 *
 * @param options that will be passed into the mutation, supported options are listed on: https://v4.apollo.vuejs.org/guide-composable/mutation.html#options;
 *
 * @example
 * const { mutate, loading, error, onDone } = useNewBookmarkMutation({
 *   variables: {
 *     name: // value for 'name'
 *     url: // value for 'url'
 *     color: // value for 'color'
 *   },
 * });
 */
export function useNewBookmarkMutation(
  options:
    | VueApolloComposable.UseMutationOptions<
        NewBookmarkMutation,
        NewBookmarkMutationVariables
      >
    | ReactiveFunction<
        VueApolloComposable.UseMutationOptions<
          NewBookmarkMutation,
          NewBookmarkMutationVariables
        >
      >
) {
  return VueApolloComposable.useMutation<
    NewBookmarkMutation,
    NewBookmarkMutationVariables
  >(NewBookmarkDocument, options);
}
export type NewBookmarkMutationCompositionFunctionResult = VueApolloComposable.UseMutationReturn<
  NewBookmarkMutation,
  NewBookmarkMutationVariables
>;
export const RegisterDocument = gql`
  mutation register($email: String!, $password: String!) {
    register(email: $email, password: $password) {
      token
      user {
        id
        email
      }
    }
  }
`;

/**
 * __useRegisterMutation__
 *
 * To run a mutation, you first call `useRegisterMutation` within a Vue component and pass it any options that fit your needs.
 * When your component renders, `useRegisterMutation` returns an object that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - Several other properties: https://v4.apollo.vuejs.org/api/use-mutation.html#return
 *
 * @param options that will be passed into the mutation, supported options are listed on: https://v4.apollo.vuejs.org/guide-composable/mutation.html#options;
 *
 * @example
 * const { mutate, loading, error, onDone } = useRegisterMutation({
 *   variables: {
 *     email: // value for 'email'
 *     password: // value for 'password'
 *   },
 * });
 */
export function useRegisterMutation(
  options:
    | VueApolloComposable.UseMutationOptions<
        RegisterMutation,
        RegisterMutationVariables
      >
    | ReactiveFunction<
        VueApolloComposable.UseMutationOptions<
          RegisterMutation,
          RegisterMutationVariables
        >
      >
) {
  return VueApolloComposable.useMutation<
    RegisterMutation,
    RegisterMutationVariables
  >(RegisterDocument, options);
}
export type RegisterMutationCompositionFunctionResult = VueApolloComposable.UseMutationReturn<
  RegisterMutation,
  RegisterMutationVariables
>;
export const UpdateBookmarkDocument = gql`
  mutation updateBookmark(
    $id: ID!
    $name: String
    $url: String
    $color: String
  ) {
    updateBookmark(id: $id, name: $name, url: $url, color: $color) {
      id
      name
      url
      color
    }
  }
`;

/**
 * __useUpdateBookmarkMutation__
 *
 * To run a mutation, you first call `useUpdateBookmarkMutation` within a Vue component and pass it any options that fit your needs.
 * When your component renders, `useUpdateBookmarkMutation` returns an object that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - Several other properties: https://v4.apollo.vuejs.org/api/use-mutation.html#return
 *
 * @param options that will be passed into the mutation, supported options are listed on: https://v4.apollo.vuejs.org/guide-composable/mutation.html#options;
 *
 * @example
 * const { mutate, loading, error, onDone } = useUpdateBookmarkMutation({
 *   variables: {
 *     id: // value for 'id'
 *     name: // value for 'name'
 *     url: // value for 'url'
 *     color: // value for 'color'
 *   },
 * });
 */
export function useUpdateBookmarkMutation(
  options:
    | VueApolloComposable.UseMutationOptions<
        UpdateBookmarkMutation,
        UpdateBookmarkMutationVariables
      >
    | ReactiveFunction<
        VueApolloComposable.UseMutationOptions<
          UpdateBookmarkMutation,
          UpdateBookmarkMutationVariables
        >
      >
) {
  return VueApolloComposable.useMutation<
    UpdateBookmarkMutation,
    UpdateBookmarkMutationVariables
  >(UpdateBookmarkDocument, options);
}
export type UpdateBookmarkMutationCompositionFunctionResult = VueApolloComposable.UseMutationReturn<
  UpdateBookmarkMutation,
  UpdateBookmarkMutationVariables
>;
export const GetBookmarksDocument = gql`
  query getBookmarks {
    bookmarks {
      id
      name
      url
      color
    }
  }
`;

/**
 * __useGetBookmarksQuery__
 *
 * To run a query within a Vue component, call `useGetBookmarksQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetBookmarksQuery` returns an object from Apollo Client that contains result, loading and error properties
 * you can use to render your UI.
 *
 * @param options that will be passed into the query, supported options are listed on: https://v4.apollo.vuejs.org/guide-composable/query.html#options;
 *
 * @example
 * const { result, loading, error } = useGetBookmarksQuery();
 */
export function useGetBookmarksQuery(
  options:
    | VueApolloComposable.UseQueryOptions<
        GetBookmarksQuery,
        GetBookmarksQueryVariables
      >
    | VueCompositionApi.Ref<
        VueApolloComposable.UseQueryOptions<
          GetBookmarksQuery,
          GetBookmarksQueryVariables
        >
      >
    | ReactiveFunction<
        VueApolloComposable.UseQueryOptions<
          GetBookmarksQuery,
          GetBookmarksQueryVariables
        >
      > = {}
) {
  return VueApolloComposable.useQuery<
    GetBookmarksQuery,
    GetBookmarksQueryVariables
  >(GetBookmarksDocument, {}, options);
}
export type GetBookmarksQueryCompositionFunctionResult = VueApolloComposable.UseQueryReturn<
  GetBookmarksQuery,
  GetBookmarksQueryVariables
>;
