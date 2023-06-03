pragma experimental ABIEncoderV2;

interface ENS {

    // Logged when the owner of a node assigns a new owner to a subnode.
    event NewOwner(bytes32 indexed node, bytes32 indexed label, address owner);

    // Logged when the owner of a node transfers ownership to a new account.
    event Transfer(bytes32 indexed node, address owner);

    // Logged when the resolver for a node changes.
    event NewResolver(bytes32 indexed node, address resolver);

    // Logged when the TTL of a node changes
    event NewTTL(bytes32 indexed node, uint64 ttl);

    // Logged when an operator is added or removed.
    event ApprovalForAll(address indexed owner, address indexed operator, bool approved);

    function setRecord(bytes32 node, address owner, address resolver, uint64 ttl) external;
    function setSubnodeRecord(bytes32 node, bytes32 label, address owner, address resolver, uint64 ttl) external;
    function setSubnodeOwner(bytes32 node, bytes32 label, address owner) external returns(bytes32);
    function setResolver(bytes32 node, address resolver) external;
    function setOwner(bytes32 node, address owner) external;
    function setTTL(bytes32 node, uint64 ttl) external;
    function setApprovalForAll(address operator, bool approved) external;
    function owner(bytes32 node) external view returns (address);
    function resolver(bytes32 node) external view returns (address);
    function ttl(bytes32 node) external view returns (uint64);
    function recordExists(bytes32 node) external view returns (bool);
    function isApprovedForAll(address owner, address operator) external view returns (bool);
}


interface BaseRegistrarImplementation  {

    function setEns(ENS _ens) external  ;

    function setNode(bytes32 _baseNode) external;

    /**
     * @dev Gets the owner of the specified token ID. Names become unowned
     *      when their registration expires.
     * @param tokenId uint256 ID of the token to query the owner of
     * @return address currently marked as the owner of the given token ID
     */
    function ownerOf(uint256 tokenId) external view returns (address); 

    function setPrimaryDomain(uint256 tokenId) external returns (bool) ;

    function getPrimaryDomainId(address user) external view returns (uint256) ;

    function getPrimaryDomainName(address user) external view returns (string memory) ;
    
    function myPrimaryDomainInfo() external view returns (uint256, string memory) ;
    function getUserFromPrimaryName(string memory name) external view returns (address);

    // Authorises a controller, who can register and renew domains.
    function addController(address controller) external ;

    // Revoke controller permission for an address.
    function removeController(address controller) external ;

    // Set the resolver for the TLD this registrar manages.
    function setResolver(address resolver) external ;

    // Returns the expiration timestamp of the specified id.
    function nameExpires(uint256 id) external view returns(uint) ;

    // Returns true iff the specified name is available for registration.
    function available(uint256 id) external view returns(bool) ;

    /**
     * @dev Register a name.
     * @param id The token ID (keccak256 of the label).
     * @param owner The address that should own the registration.
     * @param duration Duration in seconds for the registration.
     */
    function register(string calldata name, uint256 id, address owner, uint duration) external returns(uint) ;

    /**
     * @dev Register a name, without modifying the registry.
     * @param id The token ID (keccak256 of the label).
     * @param owner The address that should own the registration.
     * @param duration Duration in seconds for the registration.
     */
    function registerOnly(string calldata name, uint256 id, address owner, uint duration) external returns(uint);

    function renew(uint256 id, uint duration) external  returns(uint) ;

    /**
     * @dev Reclaim ownership of a name in ENS, if you own it in the registrar.
     */
    function reclaim(uint256 id, address owner) external ;

    function supportsInterface(bytes4 interfaceID) external view returns (bool) ;

    function getDomainName(uint256 tokenId) external view returns (string memory) ;

    function getOwnedDomains(address user) external view returns (uint256[] memory, string[] memory) ;

    /**
     * @dev See {IERC721Metadata-name}.
     */
    function name() external view returns (string memory);

    /**
     * @dev See {IERC721Metadata-symbol}.
     */
    function symbol() external view returns (string memory) ;

    /**
     * @dev See {IERC721Metadata-tokenURI}.
     */
    function tokenURI(uint256 tokenId) external view  returns (string memory) ;

    function setBaseURI(string memory _newBaseURI) external ;

    function setBaseExtension(string memory _newBaseExtension) external ;

}